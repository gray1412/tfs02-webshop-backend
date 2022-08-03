package model

import "time"

type Brand struct {
	Id   int    `json:"id"`
	Name string `json:"name" `

	ImageId int `json:"image_id"`

	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
