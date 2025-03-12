package option

import (
	"fmt"
)

type FilterExp struct {
	Field string
	Op    OpType
	Value any
}

func NewExpression(field string, op OpType, value any) Filter {
	return &FilterExp{
		Field: field,
		Op:    op,
		Value: value,
	}
}

func (f FilterExp) Query() (string, []any) {
	switch f.Value.(type) {
	case []int, []int32, []int64, []float32, []float64, []string:
		query, args := extractQueryArgs(f.Value)
		return fmt.Sprintf("%s %s %s", f.Field, f.Op, query), args
	default:
		return fmt.Sprintf("%s %s ?", f.Field, f.Op), []any{f.Value}
	}
}
