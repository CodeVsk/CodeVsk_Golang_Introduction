package dto

import "github.com/google/uuid"

type ProductDto struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Amount float32   `json:"amount"`
}