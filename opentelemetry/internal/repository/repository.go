package repository

import (
	"context"

	"github.com/CodeVsk/CodeVsk_Golang_Introduction/opentelemetry/pkg/otel"
)

type IRepository interface {
	Add(ctx context.Context)
	Get(ctx context.Context) map[int]interface{}
}

type Repository struct {
	items map[int]interface{}
}

func NewRepository() *Repository {
	return &Repository{
		items: make(map[int]interface{}),
	}
}

func (r *Repository) Add(ctx context.Context) {
	v := len(r.items) + 1

	r.items[v] = v
}

func (r *Repository) Get(ctx context.Context) map[int]interface{} {
	_, span := otel.Tracer.Start(ctx, "Repository.Get")
	defer span.End()

	return r.items
}
