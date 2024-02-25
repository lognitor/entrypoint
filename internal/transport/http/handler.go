package http

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/valyala/fasthttp"
	"log"
	"log/slog"
	"os"
)

type Handler struct {
	kafka  KafkaWriterInterface
	logger *slog.Logger
}

func NewHandler(w KafkaWriterInterface) (*Handler, error) {
	return &Handler{
		kafka:  w,
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}, nil
}

func (h *Handler) Handle(ctx *fasthttp.RequestCtx) {
	//err := h.kafka.WriteMessages(ctx, kafka.Message{
	//	Key:   []byte("test"),
	//	Value: []byte(time.Now().String()),
	//	Time:  time.Now(),
	//})
	//
	//if err != nil {
	//	h.error(ctx, err)
	//	return
	//}

	go func() {
		i := 0
		for {
			go func() {
				err := h.kafka.WriteMessages(context.Background(), kafka.Message{
					Key:   []byte("test"),
					Value: []byte("test"),
				})
				if err != nil {
					panic(err)
				}
			}()
			i++
			if i == 50000 {
				log.Println("50000 messages sent")
				break
			}
		}
	}()

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBodyString("ok")
}

func (h *Handler) error(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	h.logger.Error(fmt.Sprintf("Handler error: %s", err.Error()))
	ctx.SetBodyString("error")
}
