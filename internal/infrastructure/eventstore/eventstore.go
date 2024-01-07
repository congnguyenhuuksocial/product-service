package eventstore

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"product-service/internal/core/entities"
	"time"
)

type EventStore struct {
	db *gorm.DB
}

func NewEventStore(db *gorm.DB) *EventStore {
	return &EventStore{
		db: db,
	}
}

func (store *EventStore) AppendEvent(ctx context.Context, event entities.Event) error {
	eventData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	dbEvent := &DBEvent{
		Name:    event.EventName(),
		Data:    eventData,
		Created: event.Timestamp(),
	}

	return store.db.WithContext(ctx).Create(dbEvent).Error
}

type DBEvent struct {
	gorm.Model
	Name    string    `gorm:"column:name;type:varchar(255);not null"`
	Data    []byte    `gorm:"column:data;type:json;not null"`
	Created time.Time `gorm:"column:created;type:timestamp;not null"`
}
