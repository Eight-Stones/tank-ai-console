//nolint:predeclared // package should need to this name
package error

import (
	"errors"
)

// Message return original message from erorr.
func Message(err error) string {
	if err == nil {
		return ""
	}

	var e *Error
	if !errors.As(err, &e) {
		return err.Error()
	}
	msg := e.message
	ee := e.Unwrap()
	if ee == nil || !errors.As(ee, &e) {
		return msg
	}

	if dmsg := Message(ee); dmsg != "" {
		return dmsg
	}

	return msg
}
