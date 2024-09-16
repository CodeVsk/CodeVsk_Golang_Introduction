package main

import "fmt"

func some(number_1 int8, number_2 int8) int8 {
	return number_1 + number_2
}

func calculate(number_1, number_2 int8) (int8, int8) {
	some := number_1 + number_2
	subtraction := number_1 - number_2

	return some, subtraction
}
func main() {
	total := some(2, 3)

	fmt.Println(total)

	var f = func() string {
		return "Hello World!"
	}

	fmt.Println(f())

	totalSome, totalSubtraction := calculate(1, 4)

	fmt.Println(totalSome, totalSubtraction)

	totalSome_2, _ := calculate(1, 4)

	fmt.Println(totalSome_2)
}