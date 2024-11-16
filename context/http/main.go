package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 1)

	defer cancel()

	base_url := "https://economia.awesomeapi.com.br/last/USD-BRL,EUR-BRL,BTC-BRL"

	req, err := http.NewRequestWithContext(ctx, "GET", base_url, nil)
	if err != nil {
		fmt.Println("An error ocurred while creating the request.")

		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("An error ocurred while doing the request.")

		return
	}

	defer res.Body.Close()

	fmt.Println("Request status:", res.Status)
}