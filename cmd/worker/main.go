package main

import (
	"errors"
	"fmt"
	"github.com/lognitor/entrypoint/configs"
	"github.com/lognitor/entrypoint/internal/database/ch"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
	"github.com/lognitor/entrypoint/internal/worker"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	p, err := parsArgs()
	if err != nil {
		log.Fatalf("error parsing args: %v", err)
	}

	// init kafka consumer
	kc, err := configs.NewKafka()
	if err != nil {
		log.Fatalf("error creating kafka config: %v", err)
	}
	consumer := kafka.NewConsumer(kc, uint(p))

	// init clickhouse config
	chHosts := strings.Split(",", os.Getenv("CLICKHOUSE_HOSTS"))
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
	w := worker.NewWorker(clickhouse, consumer)
	defer w.Close()

	log.Fatalln(w.Run())
}

func parsArgs() (partition uint64, err error) {
	args := os.Args[1:]
	if len(args) < 1 {
		err = errors.New("missing partition number")
		return 0, err
	}

	for _, arg := range args {
		s := strings.Split(arg, "=")
		if len(s) != 2 {
			err = errors.New("invalid partition number")
			return 0, err
		}

		switch s[0] {
		case "--partition", "-p":
			partition, err = strconv.ParseUint(s[1], 10, 32)
			if err != nil {
				return 0, err
			}
		default:
			return 0, fmt.Errorf("invalid argument name: %s", s[0])
		}
	}

	return partition, err
}
