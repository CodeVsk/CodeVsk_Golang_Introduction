package main

import (
	"database/internal/application/service"
	"database/internal/infra/database"
	"database/internal/infra/repository"
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {
	db, err := database.Connection()

	if(err != nil){
		panic(err)
	}

	defer db.Close()

	product_repository := repository.NewPgProductRepository(db)
	product_service := service.NewProductService(product_repository)

	product_id, _ := uuid.FromString("0378b6d6-45f2-4083-95f3-5201e6bc2507")

	p := product_service.GetById(product_id)

	fmt.Println(p)
}
