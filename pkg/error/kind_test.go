//nolint:predeclared,testpackage // package should need to this name
package error

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKind(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want Type
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: UnknownType,
		},
		{
			name: "simple error",
			args: args{
				err: errors.New("test error"),
			},
			want: UnknownType,
		},
		{
			name: "package error",
			args: args{
				err: &Error{
					kind: InternalType,
				},
			},
			want: InternalType,
		},
		{
			name: "wrapped package error",
			args: args{
				err: &Error{
					kind: InternalType,
					cause: &Error{
						kind: InternalType,
					},
				},
			},
			want: InternalType,
		},
		{
			name: "fmted package error",
			args: args{
				err: fmt.Errorf("%w: some error", &Error{
					kind: InternalType,
				}),
			},
			want: InternalType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, Kind(tt.args.err))
		})
	}
}
