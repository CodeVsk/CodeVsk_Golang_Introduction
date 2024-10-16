package contract

import (
	"github.com/codevsk/golang/simple-api/internal/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	GetById(Id uuid.UUID) (*entity.Product, error)
}