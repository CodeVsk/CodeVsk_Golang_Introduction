package repository

type ProductRepository interface {
	GetAllProducts()
}

type PgProductRepository struct{}

func NewPgProductRepository() ProductRepository {
	return &PgProductRepository{}
}

func (PgProductRepository) GetAllProducts() {}