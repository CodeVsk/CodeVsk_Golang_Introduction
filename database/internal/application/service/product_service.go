package service

import (
	"database/internal/domain/entity"
	"database/internal/domain/repository"
	"log"

	"github.com/gofrs/uuid"
)

type ProductService struct {
	product_repository repository.ProductRepository
}

func NewProductService(product_repository repository.ProductRepository) *ProductService {
	return &ProductService{product_repository: product_repository}
}

func (s *ProductService) GetById(Id uuid.UUID) *entity.Product {
	p, err := s.product_repository.GetById(Id);

	if err != nil {
		log.Fatal(err)
	}

	return p
}