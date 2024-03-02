package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"log/slog"
	"os"
)

type Handler struct {
	srv    ServiceInterface
	logger *slog.Logger
}

func NewHandler(srv ServiceInterface) (*Handler, error) {
	return &Handler{
		srv:    srv,
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}, nil
}

func (h *Handler) Handle(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPost() {
		h.error(ctx, errors.New("not post request"))
		return
	}

	body := ctx.PostBody()
	token := ctx.Request.Header.Peek("TOKEN")
	action := ctx.Request.URI().LastPathSegment()

	if string(action) == "fast" {
		go func(token []byte, body []byte) {
			if err := h.ValidateAndWrite(token, body); err != nil {
				h.logger.Error(fmt.Sprintf("Handler error: %s", err.Error()))
			}
		}(token, body)
	} else {
		if err := h.ValidateAndWrite(token, body); err != nil {
			h.error(ctx, err)
			return
		}
	}

	ctx.SetStatusCode(fasthttp.StatusCreated)
	ctx.SetBodyString("ok")
}

func (h *Handler) ValidateAndWrite(token []byte, body []byte) error {
	if err := h.srv.ValidateRequest(string(token), body); err != nil {
		return err
	}

	if err := h.srv.WriteRequest(context.Background(), string(token), body); err != nil {
		return err
	}

	return nil
}

func (h *Handler) error(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	h.logger.Error(fmt.Sprintf("Handler error: %s", err.Error()))
	ctx.SetBodyString("error")
}
