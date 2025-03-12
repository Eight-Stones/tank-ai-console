//nolint:predeclared,testpackage // package should need to this name
package error

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Wrap(t *testing.T) {
	t.Run("Wrap Unknown", func(t *testing.T) {
		wrapped := fmt.Errorf("some simple error")

		target := &Error{
			kind:    InternalType,
			cause:   wrapped,
			message: "wrapped text",
		}

		require.Equal(t, target, Wrap(wrapped, "wrapped text"))
	})

	t.Run("Wrap", func(t *testing.T) {
		wrapped := &Error{
			kind:  NotFoundType,
			cause: fmt.Errorf("some error"),
		}

		target := &Error{
			kind:    NotFoundType,
			cause:   wrapped,
			message: "wrapped text",
		}

		require.Equal(t, target, Wrap(wrapped, "wrapped text"))
	})

	t.Run("Wrapf", func(t *testing.T) {
		wrapped := &Error{
			kind:  InternalType,
			cause: fmt.Errorf("some error"),
		}

		target := &Error{
			kind:    InternalType,
			cause:   wrapped,
			message: "wrapped text param",
		}

		require.Equal(t, target, Wrapf(wrapped, "wrapped text %s", "param"))
	})
}
