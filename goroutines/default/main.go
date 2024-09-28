package main

import (
	"fmt"
	"time"
)

func print(text string) {
	for {
		fmt.Println(text)
		
		time.Sleep(time.Second)
	}
}

func main() {
	go print("Capybara!")
	print("Hello World!")
}