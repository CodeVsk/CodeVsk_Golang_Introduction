package main

import "fmt"

func invertValue(number int) int {
	return number * -1
}

func invertValuePointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 10
	fmt.Println(number, invertValue(number))

	number_pointer := 11
	fmt.Println(number_pointer)
	invertValuePointer(&number_pointer)
	fmt.Println(number_pointer)
}