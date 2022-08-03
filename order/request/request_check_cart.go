package request

import (
	"errors"
	"project-backend/util/constant"
)

type RequestCheckCart struct {
	Total          float64 `json:"total"`
	DiscountAmount float64 `json:"discount_amount"`
	TotalBill      float64 `json:"total_bill"`
	TotalWeight    string  `json:"total_weight"`
	VoucherCode    string  `json:"voucher_code"`
	Cart           []ItemCheckCart  `json:"cart"`
}
type ItemCheckCart struct { // variant
	Id         int64   `json:"id"`
	Alias      string  `json:"alias"`
	Image      string  `json:"image"`
	ProductName string  `json:"name"`
	Quantity   int64   `json:"quantity"`
	Stock      int64   `json:"stock"`
	Variant    Variant `json:"variant"`
}
type Variant struct {
	Id            int64   `json:"id"`
	Code          string  `json:"code"`
	VariantName   string  `json:"name"`
	Stock 		  int64   `json:"quantity"`
	Option1       int64   `json:"option1"`
	Option2       int64   `json:"option2"`
	Option3       int64   `json:"option3"`
	OriginalPrice float64 `json:"original_price"`
	Position      int64   `json:"position"`
	Price         float64 `json:"price"`
	Weight        string  `json:"weight"`
}

func (c RequestCheckCart) CheckCaculation(discountAmount float64) error {
	if c.Total < 0 || c.DiscountAmount < 0 ||  c.TotalBill < 0{
		return errors.New(constant.ERROR_CACULATE)
	}
	if c.DiscountAmount != discountAmount {
		return errors.New(constant.ERROR_CACULATE)
	}
	total := 0.0
	for _, item := range c.Cart {
		total += float64(item.Quantity) * item.Variant.Price
	}

	totalBill :=  total - c.DiscountAmount
	if total != c.Total {
		return errors.New(constant.ERROR_CACULATE)

	}
	if totalBill != c.TotalBill {
		return errors.New(constant.ERROR_CACULATE)
	}
	return nil
}
func CaculateDiscountAmount(total,discount,maxSaleAmount float64, unit string) float64 {
	discountAmount := 0.0
	switch unit {
	case "percent":
		discountAmount = total * discount / 100
		if discountAmount > maxSaleAmount {
			discountAmount = maxSaleAmount
		}
	default:
		discountAmount = discount
	}
	return discountAmount
}