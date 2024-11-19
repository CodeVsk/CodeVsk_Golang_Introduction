package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/codevsk/golang/sample-rabbitmq-go/internal/event"
	"github.com/codevsk/golang/sample-rabbitmq-go/pkg/rabbitmq"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	t := event.NewTicketOpenEvent("Backup", "Backup drivers")

	ticketJSON, err := json.Marshal(t)
	if err != nil {
		log.Fatal("Failed to convert open ticket into json", err)
	}

	err = rabbitmq.Publish(ch, "golang.ticket-oppened", ticketJSON)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Ticket oppenned successfully.")
}