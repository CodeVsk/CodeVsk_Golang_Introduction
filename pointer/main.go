package main

import "fmt"

func main() {
	var number int8
	var pointer *int8

	number = 10
	pointer = &number

	fmt.Println(number, pointer, *pointer)

	number = 2

	fmt.Println(number, pointer, *pointer)
}