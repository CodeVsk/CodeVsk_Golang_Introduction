package main

import "fmt"

func print_name() {
	fmt.Println("Capybara")
}

func print_hello() {
	fmt.Println("Hello world")
}

func main() {
	defer print_name()

	print_hello()
}