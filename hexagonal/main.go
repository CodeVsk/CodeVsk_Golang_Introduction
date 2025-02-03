package main

import (
	"log"

	"github.com/codevsk/golang/hexagonal/adapter/input/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.InitRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error trying to init application on port 8080", err)
		return
	}
}