package output

import (
	"github.com/codevsk/golang/hexagonal/application/domain"
)

type ProductsPort interface {
	GetProductsPort(domain.ProductsReqDomain) (*domain.ProductsDomain, error)
}