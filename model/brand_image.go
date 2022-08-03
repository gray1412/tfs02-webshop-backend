package model

import "time"

type BrandImage struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
	Alt string `json:"alt"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
