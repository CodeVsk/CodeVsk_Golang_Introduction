package main

import (
	"fmt"
	"time"
)

func main() {
	var name string = "Vsk"

	fmt.Println(name)

	full_name := "CodeVsk"

	fmt.Println(full_name)

	var (
		number int = 22
		current_time time.Time = time.Now()
	)

	fmt.Println(number)
	fmt.Println(current_time)

	const constant string = "capybara!"

	fmt.Println(constant)

	hello, world := "Hello", "World!"

	fmt.Println(hello, world)
}