package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Capybara"
	channel <- "Developer"

	message_1 := <-channel
	message_2 := <-channel

	fmt.Println(message_1)
	fmt.Println(message_2)
}