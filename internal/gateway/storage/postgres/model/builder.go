package model

import (
	"github.com/Masterminds/squirrel"

	"go-micro-service-template/pkg/database/option"
)

func AddSelectFilter(in squirrel.SelectBuilder, opt *option.Option) squirrel.SelectBuilder {
	if opt == nil {
		return in
	}

	where, args := opt.Query()
	if where == "" || len(args) == 0 {
		return in
	}

	return in.Where(where, args...)
}

func AddSelectLimit(in squirrel.SelectBuilder, opt *option.Option) squirrel.SelectBuilder {
	if opt == nil {
		return in
	}

	if opt.Limit > 0 {
		return in.Limit(opt.Limit)
	}

	return in
}

func AddSelectOffset(in squirrel.SelectBuilder, opt *option.Option) squirrel.SelectBuilder {
	if opt == nil {
		return in
	}

	if opt.Offset > 0 {
		return in.Limit(opt.Offset)
	}

	return in
}

func AddSelectOrder(in squirrel.SelectBuilder, opt *option.Option) squirrel.SelectBuilder {
	if opt == nil {
		return in
	}

	if opt.Order != nil {
		return in.OrderByClause(opt.OrderBy())
	}

	return in
}

func AddDeleteFilter(in squirrel.DeleteBuilder, opt *option.Option) squirrel.DeleteBuilder {
	if opt == nil {
		return in
	}

	where, args := opt.Query()
	if where == "" || len(args) == 0 {
		return in
	}

	return in.Where(where, args...)
}
