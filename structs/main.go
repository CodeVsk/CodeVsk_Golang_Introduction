package main

import "fmt"

type deposit struct {
	address string
	city string
	number int8
}

type inventory struct {
	deposit //"inheritance"
	quantity int8
}

type product struct {
	name      string
	price     int8
	inventory inventory
}

func main() {
	var inv inventory

	inv.quantity = 2

	inv.address = "potato street"
	inv.city = "potato land"
	inv.number = 1

	prd := product{"Moto G4", 100, inv}
	fmt.Println(prd)

	prd_2 := product{name: "Moto G3"}
	fmt.Println(prd_2)
}