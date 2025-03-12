package middleware

import (
	"context"
	"fmt"
	"net/http"

	"go-micro.dev/v4/server"

	"go-micro-service-template/internal/controller/rest/dto"
	er "go-micro-service-template/pkg/error"
)

// resolve defines code and message for response if error not nil.
func resolve(err error) (code int32) {
	switch er.Kind(err) {
	case er.UnknownType:
		code = http.StatusInternalServerError
	case er.CancelledType:
		code = http.StatusRequestTimeout
	case er.InvalidArgumentType:
		code = http.StatusBadRequest
	case er.DeadlineExceededType:
		code = http.StatusRequestTimeout
	case er.NotFoundType:
		code = http.StatusNotFound
	case er.AlreadyExistsType:
		code = http.StatusConflict
	case er.PermissionDeniedType:
		code = http.StatusUnauthorized
	case er.ResourceExhaustedType:
		code = http.StatusPreconditionFailed
	case er.FailedPreconditionType:
		code = http.StatusPreconditionRequired
	case er.AbortedType:
		code = http.StatusRequestTimeout
	case er.OutOfRangeType:
		code = http.StatusBadRequest
	case er.UnimplementedType:
		code = http.StatusNotImplemented
	case er.InternalType:
		code = http.StatusInternalServerError
	case er.UnavailableType:
		code = http.StatusServiceUnavailable
	case er.DataLossType:
		code = http.StatusUnprocessableEntity
	case er.UnauthenticatedType:
		code = http.StatusNetworkAuthenticationRequired
	default:
		return http.StatusInternalServerError
	}

	return http.StatusInternalServerError
}

// ResolverWrapper resolve error from handle function and select answer code.
func ResolverWrapper(handlerFunc server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := handlerFunc(ctx, req, rsp)
		if err != nil {
			return dto.NewError(
				fmt.Sprintf("%v:%v", req.Method(), req.Endpoint()),
				er.Cause(err).Error(),
				er.Message(err),
				resolve(err),
			)
		}
		return nil
	}
}
