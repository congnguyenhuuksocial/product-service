package entities

import (
	"encoding/json"
	"time"
)

// Event is the interface that all events should implement
type Event interface {
	EventName() string
	Timestamp() time.Time
}

// ProductCreatedEvent fires when a new product is created
type ProductCreatedEvent struct {
	ID          uint
	Name        string
	Description string
	Price       float64
	Created     time.Time
}

func (p ProductCreatedEvent) EventName() string {
	return "product.created"
}

func (p ProductCreatedEvent) Timestamp() time.Time {
	return p.Created
}

func (p ProductCreatedEvent) Marshal() string {
	out, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(out)
}
