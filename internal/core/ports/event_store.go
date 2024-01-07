package ports

import (
	"context"
	"product-service/internal/core/entities"
)

// EventStore is the interface for an event store
type EventStore interface {
	AppendEvent(ctx context.Context, event entities.Event) error
}
