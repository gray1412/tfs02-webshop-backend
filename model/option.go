package model

import "time"

type Option struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Position  int    `json:"position"`
	ProductId int    `json:"product_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
