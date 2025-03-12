package option // nolint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterSimple(t *testing.T) {
	tests := []struct {
		name  string
		input Filter
		query string
		args  []any
	}{
		{
			name:  "simple int",
			input: NewExpression("id", EQ, 12),
			query: "id = ?",
			args:  []any{12},
		},
		{
			name:  "simple int32",
			input: NewExpression("id", EQ, int32(12)),
			query: "id = ?",
			args:  []any{int32(12)},
		},
		{
			name:  "simple int64",
			input: NewExpression("id", EQ, int64(12)),
			query: "id = ?",
			args:  []any{int64(12)},
		},
		{
			name:  "simple float32",
			input: NewExpression("id", EQ, float32(12)),
			query: "id = ?",
			args:  []any{float32(12)},
		},
		{
			name:  "simple float32 with part",
			input: NewExpression("id", EQ, float32(12.21)),
			query: "id = ?",
			args:  []any{float32(12.21)},
		},
		{
			name:  "simple float64",
			input: NewExpression("id", EQ, float64(12)),
			query: "id = ?",
			args:  []any{float64(12)},
		},
		{
			name:  "simple float64 with part",
			input: NewExpression("id", EQ, float64(12.21)),
			query: "id = ?",
			args:  []any{float64(12.21)},
		},
		{
			name:  "simple string",
			input: NewExpression("id", EQ, "value"),
			query: "id = ?",
			args:  []any{"value"},
		},
		{
			name:  "simple string array",
			input: NewExpression("id", IN, []string{"v1", "v2"}),
			query: "id IN (?,?)",
			args:  []any{"v1", "v2"},
		},
		{
			name:  "simple numbers array",
			input: NewExpression("id", IN, []int{1, 2}),
			query: "id IN (?,?)",
			args:  []any{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, args := tt.input.Query()
			require.Equal(t, tt.query, query)
			require.Equal(t, tt.args, args)
		})
	}
}

func TestFilterComplex(t *testing.T) {
	tests := []struct {
		name  string
		input Filter
		query string
		args  []any
	}{
		{
			name: "and two int",
			input: NewCondition(
				AND,
				NewExpression("id", EQ, 12),
				NewExpression("id", EQ, 18),
			),
			query: "id = ? AND id = ?",
			args:  []any{12, 18},
		},
		{
			name: "and three int",
			input: NewCondition(
				AND,
				NewExpression("id", EQ, 12),
				NewExpression("id", EQ, 18),
				NewExpression("id", EQ, 21),
			),
			query: "id = ? AND id = ? AND id = ?",
			args:  []any{12, 18, 21},
		},
		{
			name: "or two int",
			input: NewCondition(
				OR,
				NewExpression("id", EQ, 12),
				NewExpression("id", EQ, 18),
			),
			query: "id = ? OR id = ?",
			args:  []any{12, 18},
		},
		{
			name: "and two int",
			input: NewCondition(
				AND,
				NewExpression("id", EQ, 12),
				NewCondition(
					OR,
					NewExpression("id", EQ, 14),
					NewExpression("id", EQ, 18),
				),
			),
			query: "id = ? AND (id = ? OR id = ?)",
			args:  []any{12, 14, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, args := tt.input.Query()
			require.Equal(t, tt.query, query)
			require.Equal(t, tt.args, args)
		})
	}
}
