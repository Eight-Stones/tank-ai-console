//nolint:predeclared // package should need to this name
package error

import (
	"errors"
)

// Kind return ErrorType of error if it can be defined.
func Kind(err error) Type {
	if err == nil {
		return UnknownType
	}

	var e *Error
	if !errors.As(err, &e) {
		return UnknownType
	}
	kind := e.kind
	ee := e.Unwrap()
	if ee == nil || !errors.As(ee, &e) {
		return kind
	}

	return Kind(ee)
}
