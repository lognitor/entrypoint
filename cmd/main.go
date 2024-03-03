package main

import (
	"github.com/lognitor/entrypoint/configs"
	"github.com/lognitor/entrypoint/internal/service"
	"github.com/lognitor/entrypoint/internal/transport/grpc"
	"github.com/lognitor/entrypoint/internal/transport/http"
	"github.com/lognitor/entrypoint/internal/transport/kafka"
	"log"
)

func main() {
	httpConfig := configs.NewHttpServer()
	grpcConfig := configs.NewGrpcServer()
	kafkaConfig, err := configs.NewKafka()
	if err != nil {
		log.Fatal(err)
	}

	kw := kafka.GetDefaultWriter(kafkaConfig)
	ks := kafka.NewKafka(kw)

	srv := service.NewService(ks)

	httpServer, err := http.NewServer(httpConfig, srv)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpcConfig, srv)

	ch := make(chan error)

	go func(server *http.Server, ch chan error) {
		ch <- server.Start()
	}(httpServer, ch)

	go func(server *grpc.Server, ch chan error) {
		ch <- server.Start()
	}(grpcServer, ch)

	log.Fatal(<-ch)
}
