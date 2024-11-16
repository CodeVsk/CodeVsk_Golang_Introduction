package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := rand.Intn(5) + 1 //Generate a random int with values between 1 and 5.
	
	ctx := context.Background()
	
	ctx, cancel := context.WithTimeout(ctx, time.Second * time.Duration(n))
	//ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	//go func() {
	//	time.Sleep(time.Second * time.Duration(n))
	//	cancel()
	//}()

	select {
		case <- ctx.Done():
			fmt.Println("Time limit exceeded.")
		case <- time.After(time.Second * 4):
			fmt.Println("Execution successfully completed.")
	}
}