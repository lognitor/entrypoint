package main

import (
	"github.com/lognitor/entrypoint/configs"
	"github.com/lognitor/entrypoint/internal/service"
	"github.com/lognitor/entrypoint/internal/transport/http"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
	"log"
)

func main() {
	httpConfig := configs.NewHttpServer()
	kafkaConfig, err := configs.NewKafka()
	if err != nil {
		log.Fatal(err)
	}

	kw := kafka.GetDefaultWriter(kafkaConfig)
	ks := kafka.NewKafka(kw)

	srv := service.NewService(ks)

	server, err := http.NewServer(httpConfig, srv)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
