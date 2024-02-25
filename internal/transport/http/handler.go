package http

import "github.com/valyala/fasthttp"

type Handler struct {
}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) Handle(ctx *fasthttp.RequestCtx) {
	ctx.SetBodyString("Hello, World!")
}
