package worker

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/goccy/go-json"
	"github.com/lognitor/entrypoint/internal/database/ch"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
	"github.com/lognitor/entrypoint/pkg/structs"
	gKafka "github.com/segmentio/kafka-go"
	"log"
	"sync"
	"time"
)

const (
	defaultDebounce  = time.Second * 3
	defaultBatchSize = 100
	defaultTableName = "log_request"
)

type Worker struct {
	clickhouse *ch.CH
	consumer   kafka.ConsumerInterface
	lastUpdate time.Time
	mu         sync.Mutex
	ctx        context.Context
	batch      driver.Batch
	processing bool // batch send in processing
}

func NewWorker(ctx context.Context, ch *ch.CH, c kafka.ConsumerInterface) *Worker {
	return &Worker{
		clickhouse: ch,
		consumer:   c,
		lastUpdate: time.Now(),
		ctx:        ctx,
	}
}

func (w *Worker) Close() error {
	return w.consumer.Close()
}

func (w *Worker) Run() error {
	w.initNewBatch()
	go w.watcher()

	for {
		m, err := w.consumer.ReadMessage(w.ctx)
		if err != nil {
			return fmt.Errorf("cannot read message: %w", err)
		}

		if err := w.handleMessage(m); err != nil {
			log.Printf("cannot handle message: %v", err)
		}
	}
}

func (w *Worker) watcher() {
	for {
		if w.processing {
			continue
		}

		d := time.Now().Sub(w.lastUpdate)
		l := w.batch.Rows()

		if d < defaultDebounce {
			log.Println("[INFO] Batch send from debounce", d, l)
			err := w.sendBatch()
			if err != nil {
				log.Printf("[ERROR] Cannot send batch: %v", err)
			}
		}

		if l >= defaultBatchSize {
			log.Println("[INFO] Batch send from size", d, l)
			err := w.sendBatch()
			if err != nil {
				log.Printf("[ERROR] Cannot send batch: %v", err)
			}
		}

		time.Sleep(time.Millisecond * 500)
	}
}

func (w *Worker) initNewBatch() {
	b, err := w.clickhouse.Conn.PrepareBatch(w.ctx, fmt.Sprintf("INSERT INTO %s", defaultTableName))
	if err != nil {
		log.Fatalf("[FATAL] Cannot prepare batch: %v", err)
	}

	w.batch = b
}

func (w *Worker) sendBatch() error {
	w.mu.Lock()
	w.processing = true
	defer func() {
		w.processing = false
	}()
	defer w.mu.Unlock()

	if err := w.batch.Send(); err != nil {
		return err
	}

	w.update()
	w.initNewBatch()
	return nil
}

func (w *Worker) update() {
	w.lastUpdate = time.Now()
}

func (w *Worker) handleMessage(m gKafka.Message) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	token := string(m.Key)
	// TODO: check exists token
	// send request to lognitor

	var l structs.Log
	if err := json.Unmarshal(m.Value, &l); err != nil {
		return err
	}

	w.update()

	return nil
}
