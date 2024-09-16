package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Execution recovered")
	}
}

func isGreaterTen(value int8) {
	defer recoverExecution()

	if value > 10 {
		fmt.Println("Value is greater 10.")

		return
	}

	panic("Value is not greater 10.")
}

func main() {
	isGreaterTen(8)
}