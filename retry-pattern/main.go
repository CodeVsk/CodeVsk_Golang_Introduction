package main

import (
	"errors"
	"math/rand"
	"time"
)

func main() {
	err := Retry(CallGateway, 3)
	if err != nil {
		println("Failed to call gateway:", err.Error())

		return
	}

	println("Successfully called gateway")
}

func CallGateway() error {
	if rand.Intn(2) == 0 {
		return nil
	}

	return errors.New("network error")
}

func Retry(fun func() error, retries int) error {
	for i := 0; i < retries; i++ {
		if err := fun(); err != nil {
			time.Sleep(time.Second * time.Duration(i))
			continue
		}
		return nil
	}

	return errors.New("max retries exceeded")
}
