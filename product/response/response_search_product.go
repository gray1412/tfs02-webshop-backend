package response

import "time"

type ResponseSearch struct{
	TotalPage int `json:"total_page"`
	TotalElement int `json:"total_elements"`
	PageIndex 	int `json:"page_index"`
	PageSize	int `json:"page_size"`
	Products []Product `json:"products"`
}
type Product struct {
	Id int `json:"id"`
	Alias string `json:"alias"`

	Name string `json:"name"`
	Price float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Quantity int	`json:"quantity"`
	ImageUrl	string 	`json:"image_url"`
	Active	int `json:"active"`
	CreatedAt	time.Time `json:"created_at"`
}
