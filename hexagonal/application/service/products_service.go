package service

import (
	"github.com/codevsk/golang/hexagonal/application/domain"
	"github.com/codevsk/golang/hexagonal/application/port/output"
)

type productsService struct{
	productsPort output.ProductsPort
}

func NewProductsService(productsPort output.ProductsPort) *productsService {
	return &productsService{productsPort}
}

func (ps *productsService) GetProducts(productsDomain domain.ProductsReqDomain) (*domain.ProductsDomain, error) {
	productsDomainResponse, err := ps.productsPort.GetProductsPort(productsDomain)

	return productsDomainResponse, err
}