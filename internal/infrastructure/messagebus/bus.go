package messagebus

import (
	"context"
	"product-service/internal/infrastructure/messagebus/messages"
)

type MessageBus interface {
	Publish(ctx context.Context, topic string, msg messages.KafkaMessage) error
	Subscribe(ctx context.Context, topic string, handler func(message messages.KafkaMessage) error) error
	Close() error
}
