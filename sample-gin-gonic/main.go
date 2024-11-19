package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id int `json:"id" required:"false"`
	Name string `json:"name" required:"true"`
	Price float32 `json:"price" required:"true"`
}

var products = []product{
	{Id: 1, Name: "Macbook", Price: 7000.91},
	{Id: 2, Name: "iPhone", Price: 1200.50},
	{Id: 3, Name: "iPad", Price: 600.25},
	{Id: 4, Name: "AirPods", Price: 150.00},
}

func main() {
	r := gin.Default()

	product_routes := r.Group("/product")

	product_routes.GET("/", func(ctx *gin.Context) {
		id := ctx.Query("id")

		product_id, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID, must be an integer"})
			return
		}
		
		for _, p := range products {
			if p.Id == product_id {
				ctx.JSON(http.StatusOK, p)
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{"message": "Product not found!"})
	})

	product_routes.POST("/create", func(ctx *gin.Context) {
		var newProduct product

		
		if err := ctx.ShouldBindJSON(&newProduct); err != nil {
			ctx.String(http.StatusBadRequest, "Body is invalid!")
			return
		}
		
		newProduct.Id = (len(products)+1)

		products = append(products, newProduct)

		ctx.JSON(http.StatusOK, gin.H{"message": "Product created succeffully!"})
	})

	r.Run(":9090")
}