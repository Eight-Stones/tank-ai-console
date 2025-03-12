package error

import (
	"errors"
	"fmt"
)

// Error custom error type with metadata.
type Error struct {
	kind    Type
	cause   error
	message string
}

// Error implements original Error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.message, e.cause)
}

// Unwrap return wrapped error.
func (e Error) Unwrap() error {
	return e.cause
}

// Is return true if input error equal.
func (e Error) Is(err error) bool {
	return errors.Is(e.cause, err)
}

// As return true if input error equal and put data into target.
func (e Error) As(target any) bool {
	return errors.As(e.cause, &target)
}
