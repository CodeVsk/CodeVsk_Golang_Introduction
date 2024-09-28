package main

import (
	"fmt"
	"sync"
	"time"
)

func print(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)

		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		print("Capybara!")
		waitGroup.Done()
	}()

	go func() {
		print("Developer!")
		waitGroup.Done()
	}()

		waitGroup.Wait()
}