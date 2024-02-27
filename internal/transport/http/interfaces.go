package http

import (
	"context"
	"time"
)

type ConfigInterface interface {
	GetPort() string
	GetAllowIp() string
}

type KafkaWriterInterface interface {
	WriteMessage(ctx context.Context, key string, message []byte, t time.Time) error
}
