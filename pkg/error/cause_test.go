package error

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCause(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "simple error",
			args: args{
				err: fmt.Errorf("some error"),
			},
			want: fmt.Errorf("some error"),
		},
		{
			name: "package error",
			args: args{
				err: &Error{
					kind:    InternalType,
					cause:   fmt.Errorf("some error"),
					message: "package wrapper",
				},
			},
			want: fmt.Errorf("some error"),
		},
		{
			name: "package wrapped error",
			args: args{
				err: &Error{
					kind: InternalType,
					cause: &Error{
						kind:    InternalType,
						cause:   fmt.Errorf("some error"),
						message: "package wrapper first",
					},
					message: "package wrapper second",
				},
			},
			want: fmt.Errorf("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, Cause(tt.args.err))
		})
	}
}
