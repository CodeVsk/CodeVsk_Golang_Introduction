package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	time.Sleep(time.Second * 3)

	cancel()

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()

	select {
		case <- ctx.Done():
			fmt.Println("Time limit exceeded.")
		case <- time.After(time.Second * 4):
			fmt.Printf("Execution succefully completed for id %d.", workerId)
	}
}