package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Amount float32 `json:"amount"`
}

func CreateProduct(name string, amount float32) (*Product, error){
	p := &Product {
		Id: uuid.New(),
		Name: name,
		Amount: amount,
	}

	err := p.validate()
	if(err != nil){
		return nil, err
	}

	return p, nil;
}

func (p *Product) validate() error {
	if p.Name == ""{
		return errors.New("O nome não pode ser vazio.")
	}

	if p.Amount < 0{
		return errors.New("O valor não pode ser menor que 0.")
	}

	return nil
}