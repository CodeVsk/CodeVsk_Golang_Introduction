package response

type ProductsClientResponse struct {
	Products []Product
	Total    int64
	Skip     int64
	Limit    int64
}

type Product struct {
	Id                 int64
	Title              string
	Description        string
	Category           string
	Price              float64
	DiscountPercentage float64
	Rating             float64
	Stock              int64
	Tags               []string
}