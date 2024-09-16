package main

import "fmt"

func calculate(number_1, number_2 int8) (some int8, subtraction int8) {
	some = number_1 + number_2
	subtraction = number_1 - number_2

	return
}

func main() {

	totalSome, totalSubtraction := calculate(1, 4)

	fmt.Println(totalSome, totalSubtraction)
}