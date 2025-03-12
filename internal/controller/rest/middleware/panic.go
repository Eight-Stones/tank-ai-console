package middleware

import (
	"context"
	"fmt"
	"net/http"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"

	"go-micro-service-template/internal/controller/rest/dto"
	er "go-micro-service-template/pkg/error"
)

func resolvePanic(req server.Request, r any) error {
	var err = dto.Error{
		ID:     fmt.Sprintf("%v:%v", req.Method(), req.Endpoint()),
		Code:   http.StatusInternalServerError,
		Cause:  "panic",
		Status: http.StatusText(http.StatusInternalServerError),
	}
	switch t := r.(type) {
	case string:
		err.Message = t
	case error:
		err.Message = t.Error()
	default:
		err.Message = er.UnknownType.New(fmt.Sprint(t)).Error()
	}
	return &err
}

func PanicWrapper(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
		defer func() {
			r := recover()
			if r != nil {
				logger.Log(logger.ErrorLevel, "Recovered from panic: %v", r)
				err = resolvePanic(req, r)
			}
		}()
		return handlerFunc(ctx, req, rsp)
	}
}
