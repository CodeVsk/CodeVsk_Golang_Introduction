package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Id    string `json:"id"`
	Name  string `json:"nome"`
	Email string `json:"email"`
}

func main() {
	u := user{Id: "abadsfsdffs", Name: "capybara", Email: "capybara@vsk.com"}

	userJSON, err := json.Marshal(u)
	if err != nil {
		log.Fatal("Erro ao converter user para json", err)
	}

	fmt.Println(userJSON)
	fmt.Println(bytes.NewBuffer(userJSON))

	var u2 user
	if err := json.Unmarshal(userJSON, &u2); err != nil {
		log.Fatal("Erro ao converter json para objeto", err)
	}

	fmt.Println(u2)
}