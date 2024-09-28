package main

import (
	"log"
	"net/http"
)

func capybara(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CAPYBARAAAAAAAAAAAA"))
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO WORLD!"))
	})

	http.HandleFunc("/capybara", capybara)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("Erro ao inicializar o serivodor http", err)
	}
}