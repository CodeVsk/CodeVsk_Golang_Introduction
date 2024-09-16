package main

import "fmt"

type product struct {
	name     string
	quantity int
}

func (p product) incrementProductQuantity() {
	p.quantity += 1
}

func (p *product) incrementProductQuantityPointer() {
	p.quantity++
}

func main() {
	product_1 := product{"coffee", 1}

	fmt.Println(product_1)

	product_1.incrementProductQuantity()

	fmt.Println(product_1)

	product_1.incrementProductQuantityPointer()

	fmt.Println(product_1)
}