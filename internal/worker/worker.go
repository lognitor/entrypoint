package worker

import (
	"context"
	"fmt"
	"github.com/lognitor/entrypoint/internal/database/ch"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
)

type Worker struct {
	clickhouse *ch.CH
	consumer   kafka.ConsumerInterface
	batch      map[string]Batch //TODO: add remove batch
}

func NewWorker(ch *ch.CH, c kafka.ConsumerInterface) *Worker {
	return &Worker{
		clickhouse: ch,
		consumer:   c,
	}
}

func (w *Worker) Close() error {
	return w.consumer.Close()
}

func (w *Worker) Run() error {
	for {
		m, err := w.consumer.ReadMessage(context.Background())
		if err != nil {
			return fmt.Errorf("cannot read message: %w", err)
		}

		fmt.Printf("[%d] message at offset %d: %s = %s\n", m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
