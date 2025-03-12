//nolint:predeclared,testpackage // package should need to this name
package error

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorType_New(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		target := &Error{
			kind:  InternalType,
			cause: fmt.Errorf("some error"),
		}
		require.Equal(t, target, InternalType.New("some error"))
	})

	t.Run("newf", func(t *testing.T) {
		target := &Error{
			kind:  InternalType,
			cause: fmt.Errorf("some error test"),
		}
		require.Equal(t, target, InternalType.Newf("some error %s", "test"))
	})
}

func TestErrorType_Wrap(t *testing.T) {
	t.Run("wrap", func(t *testing.T) {
		wrapped := &Error{
			kind:  InternalType,
			cause: fmt.Errorf("some error"),
		}

		target := &Error{
			kind:    InternalType,
			cause:   wrapped,
			message: "wrapped error",
		}

		require.Equal(t, target, InternalType.Wrap(wrapped, "wrapped error"))
	})

	t.Run("wrapf", func(t *testing.T) {
		wrapped := &Error{
			kind:  InternalType,
			cause: fmt.Errorf("some error"),
		}

		target := &Error{
			kind:    InternalType,
			cause:   wrapped,
			message: "wrapped error param",
		}

		require.Equal(t, target, InternalType.Wrapf(wrapped, "wrapped error %s", "param"))
	})
}
