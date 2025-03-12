package option

import (
	"fmt"
	"strings"
)

type FilterCond struct {
	OpType OpType
	Exps   []Filter
}

func NewCondition(op OpType, exps ...Filter) Filter {
	return &FilterCond{
		OpType: op,
		Exps:   exps,
	}
}

func (f FilterCond) Query() (string, []any) {
	var (
		builder strings.Builder
		args    []any
	)

	for idx, exp := range f.Exps {
		str, arg := exp.Query()
		if _, ok := exp.(*FilterCond); ok {
			str = fmt.Sprintf("(%s)", str)
		}

		builder.WriteString(str)

		if idx != len(f.Exps)-1 {
			builder.WriteString(fmt.Sprintf(" %s ", f.OpType.String()))
		}

		args = append(args, arg...)
	}

	return builder.String(), args
}
