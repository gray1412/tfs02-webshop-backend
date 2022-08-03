package model

import "time"

type ProductImage struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	Alt      string `json:"alt"`
	Position int    `json:"position"`

	ProductId int `json:"product_id"`
	VariantId int `json:"variant_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
