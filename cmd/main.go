package main

import (
	"github.com/lognitor/entrypoint/configs"
	"github.com/lognitor/entrypoint/internal/transport/http"
)

func main() {
	config := configs.NewHttpServer()
	server, err := http.NewServer(config)
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}
