package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	products := map[string] string {
		"name": "Moto G4",
		"model": "Motorola",
	}

	for index, product := range products {
		fmt.Println(index, product)
	}
	
}