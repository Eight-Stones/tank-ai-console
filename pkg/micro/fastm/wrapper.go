package fastm

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
)

func RequestIDAcquirer(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		return handlerFunc(context.WithValue(ctx, ctxRequestID, uuid.New().String()), req, rsp)
	}
}

func LogWrapper(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		l := logger.Fields(map[string]any{
			"request_method": req.Method(),
			"request_uri":    req.Endpoint(),
			"content_type":   req.ContentType(),
			"request_id":     ctx.Value(ctxRequestID).(string),
		})

		l.Log(logger.InfoLevel, "Request started")

		start := time.Now()
		err := handlerFunc(ctx, req, rsp)

		l.Fields(map[string]any{
			"response_time": time.Since(start).String(),
		}).Log(logger.InfoLevel, "Request finished")

		return err
	}
}

func PanicWrapper(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Log(logger.ErrorLevel, "Recovered from panic: %v", err)
			}
		}()
		return handlerFunc(ctx, req, rsp)
	}
}
