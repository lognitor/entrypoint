package worker

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/goccy/go-json"
	"github.com/lognitor/entrypoint/internal/database/ch"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
	"github.com/lognitor/entrypoint/pkg/structs"
	gKafka "github.com/segmentio/kafka-go"
	"log"
	"strings"
	"time"
)

const (
	defaultDebounce = time.Second
)

type Worker struct {
	clickhouse *ch.CH
	consumer   kafka.ConsumerInterface
	batch      map[string]*Batch //TODO: add remove batch
}

func NewWorker(ch *ch.CH, c kafka.ConsumerInterface) *Worker {
	return &Worker{
		clickhouse: ch,
		consumer:   c,
		batch:      make(map[string]*Batch),
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

		if err := w.handleMessage(m); err != nil {
			log.Printf("cannot handle message: %v", err)
		}
	}
}

func (w *Worker) handleMessage(m gKafka.Message) error {
	token := string(m.Key)

	v, ok := w.batch[token]
	if !ok {
		b, c, err := CreateBatch(token, defaultDebounce, 20)
		if err != nil {
			return err
		}
		go w.handleChannel(c, token)
		w.batch[token] = b
		v = b
	}

	var l structs.Log
	if err := json.Unmarshal(m.Value, &l); err != nil {
		return err
	}

	v.Add(l)
	return nil
}

func (w *Worker) handleChannel(ch chan []structs.Log, token string) {
	logs, ok := <-ch
	if !ok {
		// TODO: close batch worker and delete batch
		return
	}

	table := strings.ReplaceAll(token, "/", "_")
	table = strings.ReplaceAll(token, "-", "_")

	_, err := w.clickhouse.Conn.Query(context.Background(), fmt.Sprintf(`
create table if not exists %s
(
    id         UUID default generateUUIDv4(),
    type ENUM('text/plain' = 1, 'text/json' = 2, 'binary/download' = 3, 'binary/view' = 4),
    level ENUM('DEBUG' = 0, 'INFO' = 1, 'WARN' = 2, 'ERROR' = 3, 'FATAL' = 4, 'PANIC' = 5),
    prefix     String,
    ip Nullable(IPv4),
    agent Nullable(String),
    message    String,
    trace      Array(Map(String, String)),
    source Nullable(String),
    created_at DATETIME
)
    engine = MergeTree()
        PRIMARY KEY id;
`, table)) //TODO: replace token

	if err != nil {
		log.Printf("cannot create table: %v", err)
	}

	q := goqu.Insert(table).
		Cols("type",
			"level",
			"prefix",
			"ip",
			"agent",
			"message",
			"trace",
			"source",
			"created_at")

	for _, l := range logs {
		q.Rows(goqu.Vals{
			l.Level,
			l.Prefix,
			l.IP,
			l.Agent,
		})
	}

	s, p, e := q.ToSQL()
	log.Println(s, p, e)

	w.handleChannel(ch, token)
}
