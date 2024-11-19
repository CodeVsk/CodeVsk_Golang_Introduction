package service

import (
	contract "github.com/codevsk/golang/sample-http-api/internal/contract/repository"
	"github.com/codevsk/golang/sample-http-api/internal/dto"
	"github.com/google/uuid"
)

type ProductService struct {
	ProductRepository contract.ProductRepository
}

func NewProductService(pr contract.ProductRepository) *ProductService {
	
	return &ProductService{
		ProductRepository: pr,
	}
}

func (ps *ProductService) GetById(id uuid.UUID) (*dto.ProductDto, error) {
	p, err := ps.ProductRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	productDto := dto.ProductDto {
		Id: p.Id,
		Name: p.Name,
		Amount: p.Amount,
	}

	return &productDto, nil
}