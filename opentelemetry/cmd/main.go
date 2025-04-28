package main

import (
	"context"
	"log"
	"net/http"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/handler"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/middleware"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/repository"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/service"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	tracer, err := otel.NewTracer("fresh-market").UseJaeger("http://jaeger:14268/api/traces").Init()
	if err != nil {
		log.Fatal(err)
	}
	defer tracer.Shutdown(context.Background())

	meter, err := otel.NewMeter("fresh-market").UsePrometheus().Init()
	if err != nil {
		log.Fatal(err)
	}
	defer meter.Shutdown(context.Background())

	if err := middleware.InitMetrics(); err != nil {
		log.Fatal(err)
	}

	r := repository.NewRepository()
	s := service.NewService(r)
	h := handler.NewHandler(s)

	ch := chi.NewRouter()
	ch.Use(middleware.HTTPRequestCounter)
	ch.Get("/test", h.Get)
	ch.Post("/", h.Add)
	ch.Get("/metrics", promhttp.Handler().ServeHTTP)

	if err := http.ListenAndServe(":3030", ch); err != nil {
		panic(err)
	}
}
