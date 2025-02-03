package input

import (
	"github.com/codevsk/golang/hexagonal/application/domain"
)

type ProductsUseCase interface {
	GetProducts(domain.ProductsReqDomain) (*domain.ProductsDomain, error)
}