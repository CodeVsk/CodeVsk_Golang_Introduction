package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
	1 byte -> 8 bits
	int64 -> 64 bits 
	64/8 -> 8 bytes
	*/

	var number_1 int64 = 1000000000000000000

	fmt.Println(number_1)

	var number_2 uint64 = 10000000000000000

	fmt.Println(number_2)

	//alias
	var number_3 rune = -1000 //int32

	fmt.Println(number_3)

	var number_4 byte = 8 //int8

	fmt.Println(number_4)

	var number_5 float32 = 1.7
	var number_6 float64 = -1.7

	fmt.Println(number_5)
	fmt.Println(number_6)

	text_char := 'A'

	fmt.Println(text_char)

	var boolean bool = true

	fmt.Println(boolean)

	var err error

	fmt.Println(err)

	var err_2 error = errors.New("ERROR!")

	fmt.Println(err_2)
}