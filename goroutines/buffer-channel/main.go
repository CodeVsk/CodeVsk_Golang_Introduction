package main

import "fmt"

//func main() {
//	channel := make(chan string, 2)
//
//	channel <- "Capybara"
//	channel <- "Developer"
//
//	message_1 := <-channel
//	message_2 := <-channel
//
//	fmt.Println(message_1)
//	fmt.Println(message_2)
//}


func numbers(cn chan <- int ) {
	for i:=0; i < 15; i++ {
		fmt.Println("Write in channel:", i)
		cn <- i
	}

	close(cn)
}

func main() {
	cn := make(chan int, 5)

	go numbers(cn)

	for number := range cn {
		fmt.Println("Read from channel:", number)
	}

}