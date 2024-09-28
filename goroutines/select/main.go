package main

import (
	"fmt"
	"time"
)

func main() {
	channel_1, channel_2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			channel_1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel_2 <- "Channel 2"
		}
	}()

	for {
		select {
			case message_1 := <-channel_1:
				fmt.Println(message_1)
			case message_2 := <-channel_2:
				fmt.Println(message_2)
		}
	}
}