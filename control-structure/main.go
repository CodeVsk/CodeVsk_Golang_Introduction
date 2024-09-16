package main

import "fmt"

func main() {
	number := 11

	if number > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is less than 10")
	}

	if otherNumer := 5; otherNumer > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is less than 10")
	}
}