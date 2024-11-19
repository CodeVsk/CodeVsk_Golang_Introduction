package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/codevsk/golang/sample-rabbitmq-go/internal/event"
	"github.com/codevsk/golang/sample-rabbitmq-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	messages := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, "golang.ticket-oppened", messages)
	
	for message := range messages {
		var t event.TicketOpenEvent
		if err := json.Unmarshal(message.Body, &t); err != nil {
			log.Fatal("Failed to convert json into open ticket event", err)
		}

		fmt.Println("Id: ", t.Id)
		fmt.Println("Title: ", t.Title)
		fmt.Println("Description: ", t.Description)

		fmt.Println("Ticket processed successfully.")

		message.Ack(false)
	}

}