package model

import "time"

type OptionValue struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	OptionId int    `json:"option_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
