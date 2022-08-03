package model

import "time"

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Alias string `json:"alias"`

	ImageUrl      string  `json:"image_url"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Quantity      int     `json:"quantity"`
	Description   string  `json:"description"`

	BrandId    int    `json:"brand_id"`
	BrandName  string `json:"brand_name"`
	CategoryId int    `json:"category_id"`

	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
