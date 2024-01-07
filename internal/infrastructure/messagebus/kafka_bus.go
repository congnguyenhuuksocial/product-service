package messagebus

import (
	"context"
	"github.com/segmentio/kafka-go"
	"product-service/internal/infrastructure/messagebus/messages"
	"product-service/pkg/config"
)

type KafkaBus struct {
	writer *kafka.Writer
	reader *kafka.Reader
}

func NewKafkaBus(conf *config.Config) *KafkaBus {
	// Kafka writer setup
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{conf.Kafka.URI},
		Topic:   conf.Kafka.Topic,
	})

	// Kafka reader setup
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{conf.Kafka.URI},
		Topic:   conf.Kafka.Topic,
		GroupID: conf.Kafka.GroupID,
	})
	return &KafkaBus{
		writer: writer,
		reader: reader,
	}
}

func (b *KafkaBus) Publish(ctx context.Context, topic string, msg messages.KafkaMessage) error {
	return b.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(msg.Key),
		Value: []byte(msg.Value),
	})
}

func (b *KafkaBus) Subscribe(ctx context.Context, topic string, handler func(message messages.KafkaMessage) error) error {
	for {
		m, err := b.reader.ReadMessage(ctx)
		if err != nil {
			return err // or handle differently, e.g., log and continue
		}

		if err := handler(messages.KafkaMessage{Key: string(m.Key), Value: string(m.Value)}); err != nil {
			// handle handler error
		}
	}
}

func (b *KafkaBus) Close() error {
	if err := b.writer.Close(); err != nil {
		return err
	}
	return b.reader.Close()
}
