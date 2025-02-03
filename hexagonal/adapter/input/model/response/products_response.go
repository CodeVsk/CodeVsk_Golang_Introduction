package response

type ProductsResponse struct {
	Products []Product `json:"products"`
	Total    int64     `json:"total"`
	Skip     int64     `json:"skip"`
	Limit    int64     `json:"limit"`
}

type Product struct {
	Id                 int64    `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Category           string   `json:"category"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discount_percentage"`
	Rating             float64  `json:"rating"`
	Stock              int64    `json:"stock"`
	Tags               []string `json:"tags"`
}