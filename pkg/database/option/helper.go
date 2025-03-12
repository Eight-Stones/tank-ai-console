package option

import (
	"strings"

	"golang.org/x/exp/constraints"
)

type Slice interface {
	constraints.Integer | constraints.Float | string
}

func valueQueryArgs[T Slice](in []T) (string, []any) {
	var (
		builder = strings.Builder{}
		args    = make([]any, len(in))
	)

	builder.WriteRune('(')
	for idx, v := range in {
		args[idx] = v
		builder.WriteString("?")
		if idx < len(in)-1 {
			builder.WriteRune(',')
		}
	}
	builder.WriteRune(')')

	return builder.String(), args
}

func extractQueryArgs(in any) (string, []any) {
	switch in.(type) {
	case []int64:
		return valueQueryArgs[int64](in.([]int64)) //nolint
	case []float64:
		return valueQueryArgs[float64](in.([]float64)) //nolint
	case []int32:
		return valueQueryArgs[int32](in.([]int32)) //nolint
	case []int:
		return valueQueryArgs[int](in.([]int)) //nolint
	case []string:
		return valueQueryArgs[string](in.([]string)) //nolint
	}

	return "", nil
}
