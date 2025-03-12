//nolint:predeclared // package should need to this name
package error

import (
	"errors"
)

// Cause return cause error from all chain.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	var e *Error
	if errors.As(err, &e) {
		return Cause(e.Unwrap())
	}

	return err
}
