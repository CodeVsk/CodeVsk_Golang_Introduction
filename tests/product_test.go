package tests

import "testing"

func TestCreateProduct(t *testing.T) {
	prd, err := CreateProduct("Moto G4", 100, "Motorola")

	if err != nil {
		t.Error("Ocorreu um erro na criação do produto", err)
	}

	if prd.model != "Motorola" {
		t.Error("O modelo do produto está errado")
	}
}