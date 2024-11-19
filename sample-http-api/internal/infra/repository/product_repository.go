package repository

import (
	"database/sql"

	"github.com/codevsk/golang/sample-http-api/internal/entity"
	"github.com/google/uuid"
)

type DbProductRepository struct {
	DB *sql.DB
}

func NewDbProductRepository(db *sql.DB) *DbProductRepository {
	return &DbProductRepository{DB: db}
} 

func (r *DbProductRepository) GetById(id uuid.UUID) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.DB.QueryRow(`SELECT "Id", "Name", "Amount" FROM public."Products" WHERE "Id" = $1`, id).Scan(&p.Id, &p.Name, &p.Amount)

	if(err != nil){
		return nil, err
	}

	return p, nil
}