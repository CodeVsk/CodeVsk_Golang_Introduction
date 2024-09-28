package tests

import (
	"errors"

	"github.com/gofrs/uuid"
)

type product struct {
	id    uuid.UUID
	name  string
	amout float32
	model string
}

func CreateProduct(name string, amout float32, model string) (*product, error) {
	id := uuid.Must(uuid.NewV4())

	prd := &product{id, name, amout, model}

	if err := valid(prd); err != nil {
		return nil, err
	}

	return prd, nil
}

func valid(p *product) error {
	if p.name == "" || p.model == "" || p.amout == 0 {
		return errors.New("Um ou mais campos são inválidos")
	}
	return nil
}