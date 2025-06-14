package service

import (
	"context"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/internal/repository"
	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
)

type IService interface {
	Add(ctx context.Context)
	Get(ctx context.Context) map[int]interface{}
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) Add(ctx context.Context) {
	s.Repository.Add(ctx)
}

func (s *Service) Get(ctx context.Context) map[int]interface{} {
	ctx, span := otel.Tracer.Start(ctx, "Service.Get")
	defer span.End()

	return s.Repository.Get(ctx)
}
