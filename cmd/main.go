package main

import (
	"github.com/lognitor/entrypoint/configs"
	"github.com/lognitor/entrypoint/internal/transport/http"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
)

func main() {
	httpConfig := configs.NewHttpServer()
	kafkaConfig, err := configs.NewKafka()
	if err != nil {
		panic(err)
	}

	kw := kafka.GetDefaultWriter(kafkaConfig)
	ks := kafka.NewKafka(kw)

	server, err := http.NewServer(httpConfig, ks)
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
