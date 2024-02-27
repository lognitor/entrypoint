package http

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log/slog"
	"os"
	"runtime/debug"
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
	stk := debug.Stack()
	fmt.Println(stk)
	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBodyString("ok")
}

func (h *Handler) error(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	h.logger.Error(fmt.Sprintf("Handler error: %s", err.Error()))
	ctx.SetBodyString("error")
}
