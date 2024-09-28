package main

import (
	"fmt"
	"time"
)

func main() {
	channel := print("Capybara")

	for i:=0; i<5;i++ { 
		fmt.Println(<-channel)
	}
}

func print(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- text

			time.Sleep(time.Second)
		}
	}()

	return channel
}