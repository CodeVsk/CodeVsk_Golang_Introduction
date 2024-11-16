package event

import (
	"github.com/google/uuid"
)

type TicketOpenEvent struct {
	Id          uuid.UUID `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewTicketOpenEvent(title string, description string) (*TicketOpenEvent) {
	return &TicketOpenEvent{
		Id: uuid.New(),
		Title: title,
		Description: description,
	}
}