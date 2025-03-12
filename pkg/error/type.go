//nolint:predeclared // package should need to this name
package error

import (
	"fmt"
)

// Type type alias for error types.
type Type uint32

const (
	UnknownType Type = iota
	CancelledType
	InvalidArgumentType
	DeadlineExceededType
	NotFoundType
	AlreadyExistsType
	PermissionDeniedType
	ResourceExhaustedType
	FailedPreconditionType
	AbortedType
	OutOfRangeType
	UnimplementedType
	InternalType
	UnavailableType
	DataLossType
	UnauthenticatedType
)

// String implements original String interface.
func (e Type) String() string {
	switch e {
	case UnknownType:
		return "UNKNOWN"
	case CancelledType:
		return "CANCELLED"
	case InvalidArgumentType:
		return "INVALID_ARGUMENT"
	case DeadlineExceededType:
		return "DEADLINE_EXCEEDED"
	case NotFoundType:
		return "NOT_FOUND"
	case AlreadyExistsType:
		return "ALREADY_EXISTS"
	case PermissionDeniedType:
		return "PERMISSION_DENIED"
	case ResourceExhaustedType:
		return "RESOURCE_EXHAUSTED"
	case FailedPreconditionType:
		return "FAILED_PRECONDITION"
	case AbortedType:
		return "ABORTED"
	case OutOfRangeType:
		return "OUT_OF_RANGE"
	case UnimplementedType:
		return "UNIMPLEMENTED"
	case InternalType:
		return "INTERNAL"
	case UnavailableType:
		return "UNAVAILABLE"
	case DataLossType:
		return "DATA_LOSS"
	case UnauthenticatedType:
		return "UNAUTHENTICATED"
	default:
		return "UNKNOWN"
	}
}

// New creates new package error with select type and create cause error from message by fmt.Errorf.
func (e Type) New(message string) error {
	return &Error{
		kind:  e,
		cause: fmt.Errorf(message), // nolint
	}
}

// Newf creates new package error with select type and create cause error from message by fmt.Errorf.
func (e Type) Newf(format string, args ...any) error {
	return &Error{
		kind:  e,
		cause: fmt.Errorf(format, args...),
	}
}

// Wrap put input error in package error and save message into message field.
func (e Type) Wrap(err error, message string) error {
	return &Error{
		kind:    e,
		cause:   err,
		message: message,
	}
}

// Wrapf put input error in package error and save format plus args into message field.
func (e Type) Wrapf(err error, format string, args ...any) error {
	return &Error{
		kind:    e,
		cause:   err,
		message: fmt.Sprintf(format, args...),
	}
}
