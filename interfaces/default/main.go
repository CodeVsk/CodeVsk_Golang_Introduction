package main

import "fmt"

type rate interface {
	calculateRate() float64
}

type product struct {
	name  string
	price float64
}

func (p *product) calculateRate() float64 {
	return p.price * 1.4
}

func main() {
	p := product{"potato", 1.200}

	rate := p.calculateRate()

	fmt.Println(rate)

}