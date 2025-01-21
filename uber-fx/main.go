package main

import (
	"fmt"

	"github.com/codevsk/golang/uber-fx/repository"
	"github.com/codevsk/golang/uber-fx/service"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(repository.NewPgProductRepository, service.NewProductService), 
		fx.Invoke(func (service *service.ProductService)  {
			fmt.Println("All dependencies injected.")
	})).Run()
}