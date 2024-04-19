package main

import (
	"context"
	"github.com/lognitor/entrypoint/configs"
	"github.com/lognitor/entrypoint/internal/database/ch"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
	"github.com/lognitor/entrypoint/internal/worker"
	"log"
	"os"
	"strings"
)

func main() {
	// init kafka consumer
	kc, err := configs.NewKafka()
	if err != nil {
		log.Fatalf("error creating kafka config: %v", err)
	}
	consumer := kafka.NewConsumer(kc)

	// init clickhouse config
	chHosts := strings.Split(os.Getenv("CLICKHOUSE_HOSTS"), ",")
	chConfig, err := configs.NewClickHouseConfig(chHosts...)
	if err != nil {
		log.Fatalf("error creating clickhouse config: %v", err)
	}

	// init clickhouse connection
	clickhouse, err := ch.NewClickHouse(chConfig)
	if err != nil {
		log.Fatalf("error creating clickhouse service: %v", err)
	}
	defer clickhouse.Conn.Close()

	// init worker
	w := worker.NewWorker(context.TODO(), clickhouse, consumer)
	defer w.Close()

	log.Fatalln(w.Run())
}
