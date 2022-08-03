package response

type ResponseProductByID struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name"`
	Alias         string    `json:"alias"`
	Price         float64   `json:"price"`
	OriginalPrice float64   `json:"original_price"`
	Quantity      int64     `json:"quantity"`
	ImageUrl      string    `json:"image_url"`
	Description   string    `json:"description"`
	CategoryId    int64     `json:"category_id"`
	BrandName     string    `json:"brand_name"`
	Active        int64     `json:"active"`
	Options       []Option  `json:"options"`
	Variants      []Variant `json:"variants"`
	Images        []Image   `json:"images"`
}
type Option struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Position     int64         `json:"position"`
	OptionValues []OptionValue `json:"option_values"`
}
type OptionValue struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"position"`
	OptionId int64  `json:"optionsId"`
}
type Variant struct {
	Id            int64   `json:"id"`
	Code          string  `json:"code"`
	Quantity      int64     `json:"quantity"`
	Name          string  `json:"name"`
	Option1       int64   `json:"option1"`
	Option2       int64   `json:"option2"`
	Option3       int64   `json:"option3"`
	Position      int64   `json:"position"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Weight        string  `json:"weight"`
}
type Image struct {
	Id        int64  `json:"id"`
	Url       string `json:"url"`
	Alt       string `json:"alt"`
	Position  int64  `json:"position"`
	VariantId int64  `json:"variant_id"`
}

func (res *ResponseProductByID) SetOptions(opv *[]Option) {
	res.Options = make([]Option, len(*opv))
	copy(res.Options, *opv)
}
func (res *ResponseProductByID) SetVariants(v *[]Variant) {
	res.Variants = make([]Variant, len(*v))
	copy(res.Variants, *v)
}
func (res *ResponseProductByID) SetImages(i *[]Image) {
	res.Images = make([]Image, len(*i))
	copy(res.Images, *i)
}
