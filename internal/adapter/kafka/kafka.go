package kafka

import (
	"github.com/segmentio/kafka-go"
	"product-service/pkg/config"
)

type Kafka struct {
	conf *config.Config
}

type IKafka interface {
	Read() *kafka.Reader
	Write() *kafka.Writer
}

func NewKafka(conf *config.Config) *Kafka {
	return &Kafka{conf: conf}
}

func (k *Kafka) Read() *kafka.Reader {
	// Kafka reader setup
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{k.conf.Kafka.URI},
		Topic:   k.conf.Kafka.Topic,
		GroupID: k.conf.Kafka.GroupID,
	})
	defer reader.Close()
	return reader
}

func (k *Kafka) Write() *kafka.Writer {
	// Kafka writer setup
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{k.conf.Kafka.URI},
		Topic:   k.conf.Kafka.Topic,
	})
	defer writer.Close()
	return writer
}
