package main

import (
	"encoding/json"
	"net/http"

	"github.com/codevsk/golang/simple-api/internal/infra/database"
	"github.com/codevsk/golang/simple-api/internal/infra/repository"
	"github.com/codevsk/golang/simple-api/internal/service"
	"github.com/google/uuid"
)

func main() {
	conn, err := database.Connection()
	if err != nil {
		panic("Database error.")
	}

	pr := repository.NewDbProductRepository(conn)

	ps := service.NewProductService(pr)

	http.HandleFunc("/product/get-by-id", func(w http.ResponseWriter, r *http.Request) {
		p, err := ps.GetById(uuid.MustParse("0378b6d6-45f2-4083-95f3-5201e6bc2507"))
		if(err != nil) {
			w.WriteHeader(400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(p)
	})

	http.ListenAndServe(":8080", nil)
}