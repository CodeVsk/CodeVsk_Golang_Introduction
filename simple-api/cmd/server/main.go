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

	defer conn.Close()

	pr := repository.NewDbProductRepository(conn)

	ps := service.NewProductService(pr)

	http.HandleFunc("/product/get-by-id", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("id")
		p, err := ps.GetById(uuid.MustParse(q))
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