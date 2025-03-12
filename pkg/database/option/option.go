package option

import (
	"fmt"
	"strings"
)

var _ Filter = &FilterCond{}
var _ Filter = &FilterExp{}

type OpType string

const (
	EQ    = "eq"
	NotEQ = "neq"
	IN    = "in"
	NIN   = "not in"
	AND   = "and"
	OR    = "or"
)

func (o OpType) String() string {
	switch o {
	case EQ:
		return "="
	case NotEQ:
		return "!="
	case IN:
		return "IN"
	case NIN:
		return "NOT IN"
	case AND:
		return "AND"
	case OR:
		return "OR"
	}
	return ""
}

type Order struct {
	Fields    []string
	Direction string
}

type Option struct {
	Select []string
	Filter Filter
	Order  *Order
	Limit  uint64
	Offset uint64
}

func (o *Option) Fields() []string {
	return o.Select
}

func (o *Option) OrderBy() string {
	if o.Order == nil {
		return "id desc"
	}
	return fmt.Sprintf("%s %s", strings.Join(o.Order.Fields, ","), o.Order.Direction)
}

func (o *Option) Query() (string, []any) {
	return o.Filter.Query()
}
