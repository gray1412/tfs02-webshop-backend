package request

import (
	"errors"
	"project-backend/util/constant"
	"project-backend/util/validator"
)

type RequestCreateOrder struct {
	// Code
	Name           string          `json:"name"`
	Phone          string          `json:"phone"`
	Address        string          `json:"address"`
	Email          string          `json:"email"`
	Note           string          `json:"note"`
	Total          float64         `json:"total"`
	DiscountAmount float64         `json:"discount_amount"`
	Shipping       float64         `json:"shipping"`
	TotalBill      float64         `json:"total_bill"`
	TotalWeight    string          `json:"total_weight"`
	VoucherCode    string          `json:"voucher_code"`
	PaymentMethod  string          `json:"payment_method"`
	Cart           []ItemCheckCart `json:"cart"`
}

func (c RequestCreateOrder) CheckCaculation(discountAmount float64) error {
	//discount := v.Discount
	//uint := v.Unit
	//maxSaleAmount := v.MaxSaleAmount
	if c.Total < 0 || c.DiscountAmount < 0 || c.Shipping < 0 || c.TotalBill < 0 {
		return errors.New(constant.ERROR_CACULATE)
	}
	if c.DiscountAmount != discountAmount {
		return errors.New(constant.ERROR_CACULATE)
	}
	total := 0.0
	totalBill := c.Shipping - c.DiscountAmount
	for _, item := range c.Cart {
		total += float64(item.Quantity) * item.Variant.Price
	}

	totalBill += total
	if total != c.Total {
		return errors.New(constant.ERROR_CACULATE)
	}
	if totalBill != c.TotalBill {
		return errors.New("4"+constant.ERROR_CACULATE)
	}
	return nil
}

// check total weight?

func (c RequestCreateOrder) ValidCustomerInformation() error {
	// check name
	if !(validator.CheckLength(c.Name, 50) && validator.CheckName(c.Name)) {
		return errors.New(constant.ERROR_BAD_REQUEST)
	}
	if !validator.CheckPhone(c.Phone) {
		return errors.New(constant.ERROR_BAD_REQUEST)
	}
	if !validator.CheckLength(c.Address, 255) {
		return errors.New(constant.ERROR_BAD_REQUEST)
	}
	if !validator.CheckMail(c.Email) {
		return errors.New(constant.ERROR_BAD_REQUEST)
	}
	if !validator.CheckLength(c.Note, 255) {
		return errors.New(constant.ERROR_BAD_REQUEST)
	}
	return nil
}

// func (i Item) validItem() error {
// 	if !validator.CheckLength(i.ProductName, 255) {
// 		return errors.New("invalid productName ")
// 	}
// 	if !validator.CheckLength(i.VariantName, 255) {
// 		return errors.New("invalid variantName ")
// 	}

// 	if i.Price < 0 {
// 		return errors.New("invalid price ")
// 	}
// 	if i.OriginalPrice != 0 && i.Price > i.OriginalPrice {
// 		return errors.New("invalid originalPrice ")
// 	}
// 	if i.Quantity < 0 {
// 		return errors.New("invalid quantity ")
// 	}
// 	if !validator.CheckLength(i.Weight, 255) {
// 		return errors.New("invalid weight ")
// 	}
// 	return nil
// }
