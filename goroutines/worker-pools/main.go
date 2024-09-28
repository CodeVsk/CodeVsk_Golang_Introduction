package main

import "fmt"

func main() {
	position := 30

	tasks := make(chan int, position)
	results := make(chan int, position)

	go worker(tasks, results)

	for i := 0; i < position; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < position; i++ {
		result := <-results
		fmt.Println(result)
	}


}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position - 2) + fibonacci(position - 1)
}