package response

type Brand struct {
	Id    int    `json:"id"`
	Name  string `json:"name" `
	Image string `json:"image"`
	Alt   string `json:"alt"`
}
