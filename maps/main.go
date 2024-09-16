package main

import "fmt"

func main() {
	user := map[string]string{
		"name":     "capybara",
		"username": "carpinchodev",
	}

	fmt.Println(user)

	vehicle := map[string]map[string]string {
		"owner": {
			"name": "capybara",
		},
		"details": {
			"model": "car",
			"name": "fusca",
		},
	}

	fmt.Println(vehicle)

	delete(vehicle, "owner")

	fmt.Println(vehicle)
}