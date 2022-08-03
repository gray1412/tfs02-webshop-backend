package model

import (
	"project-backend/order/request"
	"time"
)

type OrderDetail struct {
	Id            int64   `json:"id"`
	ProductName   string  `json:"product_name"`
	VariantName   string  `json:"variant_name"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Quantity      int64   `json:"quantity"`
	Weight        string  `json:"weight"`

	VariantId int64 `json:"variant_id"`
	OrderId   int64 `json:"order_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (od *OrderDetail) MapFromCreateOrder(i request.ItemCheckCart) {
	od.ProductName = i.ProductName
	od.VariantName = i.Variant.VariantName
	od.Price = i.Variant.Price
	od.OriginalPrice = i.Variant.OriginalPrice
	od.Quantity = i.Quantity
	od.Weight = i.Variant.Weight
	od.VariantId = i.Variant.Id
	od.CreatedAt = time.Now()

}
