package model

import "time"

type Variant struct {
	Id            int     `json:"id"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Position      int     `json:"position"`
	Option1       int     `json:"option1"`
	Option2       int     `json:"option2"`
	Option3       int     `json:"option3"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Weight        string  `json:"weight"`
	Quantity      int     `json:"quantity"`

	ProductId int `json:"product_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
