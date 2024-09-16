package main

import "fmt"

func closure() func() {
	name := "capybara"

	function := func() {
		fmt.Println(name)
	}

	return function
}

func main() {
	name := "vsk"

	fmt.Println(name)

	function := closure()

	function()
}