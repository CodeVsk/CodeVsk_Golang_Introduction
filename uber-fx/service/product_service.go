package service

import "github.com/codevsk/golang/uber-fx/repository"

type ProductService struct {
	repository *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository) *ProductService{
	return &ProductService{
		repository: repository,
	}
}