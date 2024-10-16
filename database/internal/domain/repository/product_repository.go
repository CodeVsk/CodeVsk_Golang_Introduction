package repository

import (
	"database/internal/domain/entity"

	"github.com/gofrs/uuid"
)

type ProductRepository interface {
	GetById(Id uuid.UUID) (*entity.Product, error)
}