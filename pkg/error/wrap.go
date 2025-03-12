package error

import (
	"fmt"
)

// wrapf try to define type of input error and wraps into package error.
func wrapf(err error, format string, args ...interface{}) error {
	kind := InternalType
	if ukind := Kind(err); ukind != UnknownType {
		kind = ukind
	}

	return &Error{
		kind:    kind,
		cause:   err,
		message: fmt.Sprintf(format, args...),
	}
}

// Wrap try to define type of input error and wraps into package error.
func Wrap(err error, msg string) error {
	return wrapf(err, msg) // nolint
}

// Wrapf try to define type of input error and wraps into package error.
func Wrapf(err error, format string, args ...any) error {
	return wrapf(err, format, args...)
}
