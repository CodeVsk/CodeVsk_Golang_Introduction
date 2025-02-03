package routes

import (
	"github.com/codevsk/golang/hexagonal/adapter/input/controller"
	"github.com/codevsk/golang/hexagonal/adapter/output/gateway"
	"github.com/codevsk/golang/hexagonal/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	productGateway := gateway.NewProductsGateway()
	productService := service.NewProductsService(productGateway)
	productController := controller.NewProductController(productService)

	r.GET("/products", productController.GetProducts)
}