//nolint:predeclared,testpackage // package should need to this name
package error

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError_Error(t *testing.T) {
	type fields struct {
		kind    Type
		cause   error
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple error",
			fields: fields{
				kind:    InternalType,
				cause:   fmt.Errorf("cause error"),
				message: "message",
			},
			want: "message: cause error",
		},
		{
			name: "wrapped error",
			fields: fields{
				kind: InternalType,
				cause: Error{
					kind:    InternalType,
					cause:   fmt.Errorf("cause error"),
					message: "message",
				},
				message: "wrapped message",
			},
			want: "wrapped message: message: cause error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				kind:    tt.fields.kind,
				cause:   tt.fields.cause,
				message: tt.fields.message,
			}
			require.Equal(t, tt.want, e.Error())
		})
	}
}

func TestError_Is(t *testing.T) {
	source := &Error{kind: InternalType, cause: fmt.Errorf("cause error"), message: "cause message"}
	t.Run("check wrapped fmt", func(t *testing.T) {
		check := fmt.Errorf("wrapped: %w", source)
		require.ErrorIs(t, check, source)
	})

	t.Run("check wrapped package", func(t *testing.T) {
		check := &Error{
			kind:    InternalType,
			cause:   source,
			message: "wrapped",
		}
		require.ErrorIs(t, check, source)
	})
}

func TestError_As(t *testing.T) {
	source := &Error{kind: InternalType, cause: fmt.Errorf("cause error"), message: "cause message"}
	t.Run("check wrapped fmt", func(t *testing.T) {
		var target *Error
		require.ErrorAs(t, source, &target)
		require.Equal(t, target.message, source.message)
	})
}
