package main

import "fmt"

func main() {
	(func(name string) {
		fmt.Println(name)
	})("capybara")

	name := func(name string) string {
		return name
	}("capybara")

	fmt.Println(name)
}