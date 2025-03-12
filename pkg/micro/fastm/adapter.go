package fastm

import (
	"bytes"
	"io"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type contextAdapter struct {
	ctx *routing.Context

	r io.Reader
	w io.Writer
}

func newContextAdapter(ctx *routing.Context) *contextAdapter {
	return &contextAdapter{
		ctx: ctx,
		r:   bytes.NewReader(ctx.Request.Body()),
		w:   ctx.Response.BodyWriter(),
	}
}

func (a *contextAdapter) Read(p []byte) (int, error) {
	return a.r.Read(p)
}

func (a contextAdapter) Write(p []byte) (int, error) {
	return a.w.Write(p)
}

func (a contextAdapter) Close() error {
	return nil
}

// routingHandlerWrapper wraps basic fasthttp.RequestHandler.
func routingHandlerWrapper(h fasthttp.RequestHandler) routing.Handler {
	return func(ctx *routing.Context) error {
		fastCtx := ctx.RequestCtx
		h(fastCtx)
		return fastCtx.Err()
	}
}
