package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Email validation")
	err := checkmail.ValidateFormat("codevsk@go.com")
	fmt.Println(err)
}