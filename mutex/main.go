package main

import (
	"fmt"
	"sync"
)

func main() {
	var count uint64 
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := 0; y < 200; y++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println(count)
}