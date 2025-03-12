//nolint:predeclared,testpackage // package should need to this name
package error

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: "",
		},
		{
			name: "simple error",
			args: args{
				err: fmt.Errorf("some simple error"),
			},
			want: "some simple error",
		},
		{
			name: "package error",
			args: args{
				err: &Error{
					kind:    InternalType,
					cause:   fmt.Errorf("test error"),
					message: "text from package error",
				},
			},
			want: "text from package error",
		},
		{
			name: "wrapped package error",
			args: args{
				err: &Error{
					kind: InternalType,
					cause: &Error{
						kind:    InternalType,
						cause:   fmt.Errorf("test error"),
						message: "wrapped package error text",
					},
					message: "wrapping error text",
				},
			},
			want: "wrapped package error text",
		},
		{
			name: "double wrapped package error",
			args: args{
				err: &Error{
					kind: InternalType,
					cause: &Error{
						kind: InternalType,
						cause: &Error{
							kind:    InternalType,
							cause:   fmt.Errorf("test error"),
							message: "wrapped package error text",
						},
						message: "wrapping error text",
					},
					message: "double wrapping error text",
				},
			},
			want: "wrapped package error text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, Message(tt.args.err))
		})
	}
}
