package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"project-backend/database"
	"project-backend/model"
	"project-backend/order/request"
	"project-backend/util/constant"
	response "project-backend/util/response"
	"time"

	"github.com/gorilla/mux"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"gorm.io/gorm"
)

// var db *gorm.DB

func CheckCart(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	//	1 Get user ? Sau sửa sau khi login sẽ gửi thông tin người dùng và lưu vào state
	//	emailUser := r.Header.Get("email")
	//	user := model.User{}
	//	err := user.GetByEmail(emailUser)
	//	if err != nil {
	//		response.RespondWithJSON(w, 401, 0, "User not exists", nil)
	//		return
	//	}
	// 2. get request
	requestCart := request.RequestCheckCart{}
	err := json.NewDecoder(r.Body).Decode(&requestCart)
	if err != nil {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_BAD_REQUEST, nil)
		return
	}
	if len(requestCart.Cart) == 0 {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_BAD_REQUEST, nil)
		return
	}
	// 3. Check voucher
	discount := 0.0
	unit := constant.UNIT_USA
	maxSaleAmount := 0.0
	if requestCart.VoucherCode != "" {
		voucher := model.Voucher{}
		err = voucher.GetByCode(requestCart.VoucherCode)
		if err != nil {
			response.RespondWithJSON(w, 400, 0, constant.ERROR_VOUCHER_NOT_EXISTS, nil)
			return
		}
		if !time.Now().Before(voucher.TimeEnd) {
			response.RespondWithJSON(w, 400, 0, constant.ERROR_VOUCHER_EXPIRED, nil)
			return
		}
		discount = voucher.Discount
		unit = voucher.Unit
		maxSaleAmount = voucher.MaxSaleAmount
	}
	// check variant
	for _, item := range requestCart.Cart {
		err = checkItem(db, &item)
		if err != nil {
			response.RespondWithJSON(w, 400, 0, err.Error(), item)
			return
		}
	}
	discountAmount := request.CaculateDiscountAmount(requestCart.Total, discount, maxSaleAmount, unit)
	// check total
	err = requestCart.CheckCaculation(discountAmount)
	if err != nil {
		response.RespondWithJSON(w, 400, 0, err.Error(), nil)
		return
	}
	response.RespondWithJSON(w, 200, 1, constant.SUCCESS, nil)
}

func checkItem(db *gorm.DB, item *request.ItemCheckCart) error {
	//check price, quantity, variant exist?
	sql := "SELECT quantity FROM variants WHERE id = ?  AND price = ? AND product_id = ?"
	quantity := 0
	db.Raw(sql, item.Variant.Id, item.Variant.Price, item.Id).Scan(&quantity)
	if item.Quantity > int64(quantity) {
		return errors.New(constant.ERROR_PRODUCT_CHANGED)
	}
	return nil
}
func GetVoucherByCode(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	code := param["code"]
	voucher := model.Voucher{}
	err := voucher.GetByCode(code)

	if err != nil {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_VOUCHER_NOT_EXISTS, nil)
		return
	}
	if !time.Now().Before(voucher.TimeEnd) {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_VOUCHER_EXPIRED, nil)
		return
	}
	response.RespondWithJSON(w, 200, 1, "", voucher)

}
func CreateOrderStripe(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	//getuser
	emailUser := r.Header.Get("email")
	user := model.User{}
	err := user.GetByEmail(emailUser)
	if err != nil {
		response.RespondWithJSON(w, 401, 0, constant.ERROR_VOUCHER_NOT_EXISTS, nil)
		return
	}
	//get request
	requestStripe := request.RequestPaymentStripe{}
	err = json.NewDecoder(r.Body).Decode(&requestStripe)
	if err != nil {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_BAD_REQUEST, nil)
		return
	}
	if len(requestStripe.Order.Cart) == 0 {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_BAD_REQUEST, nil)
		return
	}
	//ktra tinh toan, luu db active: 1
	orderId, err := checkInforOrder(db, &w, requestStripe.Order, user.Id)
	if err != nil {
		return
	}
	//authen và cập nhập db active: 2
	//fmt.Println("key",requestStripe.PaymentId)
	pi, err := stripeAuth(requestStripe.PaymentId, requestStripe.Order.TotalBill, orderId)
	// fmt.Println("pi  ",err)
	if pi.Status != "succeeded" {
		response.RespondWithJSON(w, 400, 0, string(pi.Status), "")
		return
	}
	//  cập nhập db active: 2
	sql := "UPDATE orders SET active= 2 ,updated_at= ? WHERE id = ?"
	err = db.Exec(sql, time.Now(), orderId).Error
	if err != nil {
		response.RespondWithJSON(w, 500, 0, constant.ERROR_SERVER, "")
		return
	}
	response.RespondWithJSON(w, 200, 1, constant.SUCCESS, "")
}
func stripeAuth(paymentId string, total float64, orderId int64) (*stripe.PaymentIntent, error) {
	fmt.Println("key1", paymentId)
	stripe.Key = "sk_test_51JFwLUFFH821MiaEXw6A2IDjOBMmNkrPTIOikPNyzHvGDq3fPfVbjeusiryQvunwDcvbbUyryOQTc9GZf1yEwbUI00gKK04uGY"
	params := &stripe.PaymentIntentParams{
		Description:        stripe.String(string(orderId)),
		PaymentMethod:      stripe.String(paymentId),
		Amount:             stripe.Int64(int64(total * 100)),
		Currency:           stripe.String(string(stripe.CurrencyUSD)),
		Confirm:            stripe.Bool(true),
		ConfirmationMethod: stripe.String(string(stripe.PaymentIntentConfirmationMethodManual)),
	}

	return paymentintent.New(params)
}
func checkInforOrder(db *gorm.DB, w *http.ResponseWriter, requestOrder request.RequestCreateOrder, userId int) (int64, error) {
	// valid customer information
	err := requestOrder.ValidCustomerInformation()
	if err != nil {
		response.RespondWithJSON(*w, 400, 0, err.Error(), nil)
		return 0, err
	}

	//Valid voucher code
	discount := 0.0
	unit := constant.UNIT_USA
	maxSaleAmount := 0.0
	if requestOrder.VoucherCode != "" {
		voucher := model.Voucher{}
		err = voucher.CheckVoucher(requestOrder.VoucherCode)
		if err != nil {
			response.RespondWithJSON(*w, 400, 0, err.Error(), nil)
			return 0, err
		}
		discount = voucher.Discount
		unit = voucher.Unit
		maxSaleAmount = voucher.MaxSaleAmount
	}

	//check item
	for _, item := range requestOrder.Cart {
		err = checkItem(db, &item)
		if err != nil {
			response.RespondWithJSON(*w, 400, 0, err.Error(), item)
			return 0, err
		}
	}
	// check discountAmount
	discountAmount := request.CaculateDiscountAmount(requestOrder.Total, discount, maxSaleAmount, unit)
	// check total
	err = requestOrder.CheckCaculation(discountAmount)
	if err != nil {
		response.RespondWithJSON(*w, 400, 0, err.Error(), nil)
		return 0, err
	}

	// luu db
	tx := db.Begin()
	var orderDB model.Order
	orderDB.MapFromCreateOrder(&requestOrder, int64(userId))
	err = tx.Create(&orderDB).Error
	if err != nil {
		tx.Rollback()
		response.RespondWithJSON(*w, 500, 0, constant.ERROR_SERVER, nil)
		return 0, err
	}
	// update quantity
	for _, i := range requestOrder.Cart {
		sql := "UPDATE variants v SET v.quantity = v.quantity- ? WHERE v.id = ?"
		err = tx.Exec(sql, i.Quantity, i.Variant.Id).Error
		if err != nil {
			tx.Rollback()
			response.RespondWithJSON(*w, 500, 0, constant.ERROR_SERVER, i)
			return 0, err
		}
		sql = "UPDATE products p SET p.quantity = p.quantity- ? WHERE p.id = (SELECT v.product_id FROM variants  v WHERE v.id = ? )"
		err = tx.Exec(sql, i.Quantity, i.Variant.Id).Error
		if err != nil {
			tx.Rollback()
			response.RespondWithJSON(*w, 500, 0, constant.ERROR_SERVER, i)
			return 0, err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		response.RespondWithJSON(*w, 500, 0, constant.ERROR_SERVER, nil)
		return 0, err
	}
	return orderDB.Id, nil
}
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	// //1. Get user
	db := database.ConnectDB()
	emailUser := r.Header.Get("email")
	user := model.User{}
	err := user.GetByEmail(emailUser)
	if err != nil {
		response.RespondWithJSON(w, 401, 0, "User not exists", nil)
		return
	}
	// get request order
	requestOrder := request.RequestCreateOrder{}
	err = json.NewDecoder(r.Body).Decode(&requestOrder)
	if err != nil {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_BAD_REQUEST, nil)
		return
	}
	if len(requestOrder.Cart) == 0 {
		response.RespondWithJSON(w, 400, 0, constant.ERROR_BAD_REQUEST, nil)
		return
	}
	_, err = checkInforOrder(db, &w, requestOrder, user.Id)
	if err != nil {
		return
	}
	response.RespondWithJSON(w, 201, 1, constant.SUCCESS, nil)

}
