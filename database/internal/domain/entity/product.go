package entity

import "github.com/gofrs/uuid"

type Product struct {
	Id uuid.UUID
	Name string
	Amount float32
}