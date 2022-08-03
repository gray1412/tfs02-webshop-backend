package model

import (
	"project-backend/order/request"
	"time"
)

type Order struct {
	Id             int64   `json:"id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Phone          string  `json:"phone"`
	Address        string  `json:"add"`
	Email          string  `json:"email"`
	Note           string  `json:"note"`
	Total          float64 `json:"total"`
	DiscountAmount float64 `json:"discount_amount"`
	TotalBill      float64 `json:"total_bill"`
	TotalWeight    string  `json:"total_weight"`
	Shipping       float64 `json:"shipping"`

	VoucherCode        string `json:"voucher_code"`
	PaymentMethod      string `json:"payment_method"`
	SellerNote         string `json:"seller_note"`
	MailDeliveryStatus int64  `json:"mail_delivery_status"`
	UserId             int64  `json:"user_id"`

	OrderDetail []OrderDetail `json:"order_detail"`
	Active      int64         `json:"active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (o *Order) MapFromCreateOrder(r *request.RequestCreateOrder, userId int64) {
	o.Code = time.Now().String()
	o.Name = r.Name
	o.Phone = r.Phone
	o.Address = r.Address
	o.Email = r.Email
	o.Note = r.Note
	o.Total = r.Total
	o.Shipping = r.Shipping
	o.TotalBill = r.TotalBill
	o.TotalWeight = r.TotalWeight
	o.VoucherCode = r.VoucherCode
	o.PaymentMethod = r.PaymentMethod
	o.MailDeliveryStatus = 0
	o.UserId = userId
	o.Active = 1 // dat thanh cong, chua dc xu ly
	o.CreatedAt = time.Now()
	//o.OrderDetail = make([]OrderDetail, len(r.Carts))
	var od OrderDetail
	for _, i := range r.Cart {
		od.MapFromCreateOrder(i)
		o.OrderDetail = append(o.OrderDetail, od)
	}

}
