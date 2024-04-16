package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(config ConfigInterface) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        config.GetBrokers(),
		Topic:          config.GetTopic(),
		GroupID:        "workers-group",
		CommitInterval: time.Second,
		MaxBytes:       10e6, // 10MB
	})

	return &Consumer{
		reader: r,
	}
}

func (c *Consumer) ReadMessage(ctx context.Context) (kafka.Message, error) {
	return c.reader.ReadMessage(ctx)
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
