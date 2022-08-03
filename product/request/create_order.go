package request

import "time"

type Product struct {
	Id   int64    `json:"id"`
	Name string `json:"name"`
	Alias string `json:"alias"`
	ImageUrl      string  `json:"image_url"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Quantity      int64     `json:"quantity"`
	Description   string  `json:"description"`

	BrandId    int64    `json:"brand_id"`
	BrandName  string `json:"brand_name"`
	CategoryId int64    `json:"category_id"`
	Active int64 `json:"active"`
	Options []Option `json:"options"`
	Variants []Variant `json:"variants"`
	Image []ProductImage `json:"image"`
	CreatedAt time.Time `json:"created_at"`


}

type Option struct {
	Id        int64    `json:"id"`
	Name      string `json:"name"`
	Position  int64    `json:"position"`
	ProductId int64    `json:"product_id"`
	OptionValues []OptionValue `json:"option_value"`
	CreatedAt time.Time `json:"created_at"`

}
type OptionValue struct {
	Id       int64    `json:"id"`
	Name     string `json:"name"`
	Position int64    `json:"position"`
	OptionId int64    `json:"option_id"`
	CreatedAt time.Time `json:"created_at"`

}
type Variant struct {
	Id            int64     `json:"id"`
	Name          string  `json:"name"`
	Position      int64     `json:"position"`

	Option1 int64 `json:"option1"`
	Option2 int64 `json:"option2"` // req gửi lên là position của option_value của option 2
	Option3 int64 `json:"option3"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Weight        string  `json:"weight"`
	Quantity      int64     `json:"quantity"`
	ProductId int64 `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`

	//Image ProductImage `json:"image"`
}
type ProductImage struct {
	Id       int64    `json:"id"`
	Url      string `json:"url"`
	Alt      string `json:"alt"`
	Position int64    `json:"position"`

	ProductId int64 `json:"product_id"`
	VariantId int64 `json:"variant_id"` //req gui là position cua variant
	CreatedAt time.Time `json:"created_at"`

}