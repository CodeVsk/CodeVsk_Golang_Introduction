package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplex(print("Capybara"), print("Developer"))

	for i:=0; i<5;i++ { 
		fmt.Println(<-channel)
	}
}

func multiplex(channel_1, channel_2 <-chan string) <- chan string{
	channel := make(chan string)

	go func(){
		for {
			select {
				case message := <-channel_1:
					channel<-message
				case message := <-channel_2:
					channel<-message
			}
		}
	}()


	return channel
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