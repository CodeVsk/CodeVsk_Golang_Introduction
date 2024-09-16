package main

import "fmt"

func generic_print(message interface {}) {
	fmt.Println(message)
}

func main() {
	generic_print("aaaaa")
	generic_print(123)

	maps := map[interface{}]interface{} {
		"name": "capybara",
		true: 123,
	}

	fmt.Println(maps)
}