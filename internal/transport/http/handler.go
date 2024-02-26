package http

import (
	"bytes"
	"compress/gzip"
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
				str := `
{
        "Id": "102",
        "Complaints": "180",
        "Request": null,
        "JRFile": "{\"$type\":\"System.Collections.Generic.Dictionary2[[System.String, mscorlib],[System.String, mscorlib]], mscorlib\",\"PublicKey\":\"0e4bcafa-5201-4636-94e4-6e6b71d74937\",\"Name\":\"a6d5e339768690fac1ca0b5eb4c4bfb7\",\"Extension\":\"jpg\",\"MimeType\":\"image/jpeg\",\"Size\":\"37146\"}",
					"Complaint": null,
					"CreateBy": "10",
					"UpdateBy": "10",
					"CustObjVersion": "7354",
					"CreateDate": "2024-01-12T09:21:55.82624Z",
					"UpdateDate": "2024-01-12T09:21:55.82624Z"
				}
`
				var b bytes.Buffer
				gz := gzip.NewWriter(&b)
				if _, err := gz.Write([]byte(str)); err != nil {
					log.Fatal(err)
				}
				if err := gz.Close(); err != nil {
					log.Fatal(err)
				}

				err := h.kafka.WriteMessages(context.Background(), kafka.Message{
					Key:   []byte("52a44ea8-f2e0-4321-a710-fbfce7c09db9"),
					Value: b.Bytes(),
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
