package handler

import (
	"net/http"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/service"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/logger"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
	"go.uber.org/zap"
)

type Handler struct {
	Service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, span := otel.Tracer.Start(ctx, "Handler.Get")
	defer span.End()

	h.Service.Get(ctx)

	logger.Info("Get Success", zap.String("journey", "GetService"))

}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {}
