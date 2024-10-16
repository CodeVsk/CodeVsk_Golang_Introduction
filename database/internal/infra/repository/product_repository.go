package repository

import (
	"database/internal/domain/entity"
	"database/sql"

	"github.com/gofrs/uuid"
)

type PgProductRepository struct {
	DB *sql.DB
}

func NewPgProductRepository(db *sql.DB) *PgProductRepository {
	return &PgProductRepository{DB: db}
}

func (r *PgProductRepository) GetById(Id uuid.UUID) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.DB.QueryRow(`SELECT "Id", "Name", "Amount" FROM public."Products" WHERE "Id" = $1`, Id).Scan(&p.Id, &p.Name, &p.Amount)

	if(err != nil){
		return nil, err
	}

	return p, nil
}