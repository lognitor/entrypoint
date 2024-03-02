package service

import (
	"context"
	"time"
)

type KafkaWriterInterface interface {
	WriteMessage(ctx context.Context, key string, message []byte, t time.Time) error
}
