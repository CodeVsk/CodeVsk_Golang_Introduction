package request

type ProductsRequest struct {
	Type string `form:"type" binding:"required"`
}