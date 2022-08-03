package model

import "time"

type CategoryImage struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
	Alt string `json:"alt"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
