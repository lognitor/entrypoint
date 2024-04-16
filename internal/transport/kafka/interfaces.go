package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"io"
)

type ConfigInterface interface {
	GetBrokers() []string
	GetTopic() string
}

type ConsumerInterface interface {
	io.Closer
	ReadMessage(ctx context.Context) (kafka.Message, error)
}
