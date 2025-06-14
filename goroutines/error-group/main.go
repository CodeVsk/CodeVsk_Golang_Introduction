package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	err := simulateRequest()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All goroutines completed successfully")
	}
}

func simulateRequest() error {
	eg := errgroup.Group{}

	for i := 0; i < 5; i++ {
		eg.Go(func() error {
			if i == 3 {
				return fmt.Errorf("error occurred in goroutine %d", i)
			}

			return nil
		})
	}

	return eg.Wait()
}
