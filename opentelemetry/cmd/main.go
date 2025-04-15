package main

import (
	"context"
	"log"
	"net/http"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/handler"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/repository"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/service"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
	"github.com/go-chi/chi"
)

func main() {
	tracer, err := otel.NewTracer("fresh-market").UseJaeger("http://localhost:14268/api/traces").Init()
	if err != nil {
		log.Fatal(err)
	}
	defer tracer.Shutdown(context.Background())

	r := repository.NewRepository()
	s := service.NewService(r)
	h := handler.NewHandler(s)

	ch := chi.NewRouter()
	ch.Get("/", h.Get)
	ch.Post("/", h.Add)

	if err := http.ListenAndServe(":3030", ch); err != nil {
		panic(err)
	}
}
