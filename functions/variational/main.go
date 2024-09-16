package main

import "fmt"

func print_numbers(numbers... int8){
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func main() {

	print_numbers()
}