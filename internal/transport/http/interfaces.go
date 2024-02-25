package http

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type ConfigInterface interface {
	GetPort() string
	GetAllowIp() string
}

type KafkaWriterInterface interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}
