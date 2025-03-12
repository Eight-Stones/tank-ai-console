package example

import (
	"github.com/Masterminds/squirrel"

	model2 "go-micro-service-template/internal/gateway/storage/postgres/model"
	"go-micro-service-template/pkg/database/option"
)

func buildExamplesQuery(opts *option.Option) (string, []any, error) {
	builder := squirrel.
		Select(exampleFields...).
		From(exampleTable)

	builder = model2.AddSelectFilter(builder, opts)

	builder = model2.AddSelectOrder(builder, opts)

	builder = model2.AddSelectLimit(builder, opts)

	builder = model2.AddSelectOffset(builder, opts)

	return builder.PlaceholderFormat(squirrel.Dollar).ToSql()
}

func buildExamplesInsert(in *model2.Example) (string, []any, error) {
	builder := squirrel.
		Insert(exampleTable).
		Columns(exampleFields[1:]...).
		Values(in.Code, in.Name, in.Meta).
		Suffix("RETURNING id")

	return builder.PlaceholderFormat(squirrel.Dollar).ToSql()
}

func buildExamplesUpdate(in *model2.Example) (string, []any, error) {
	builder := squirrel.
		Update(exampleTable).
		Set("name", in.Name).
		Set("code", in.Code).
		Set("meta", in.Meta).
		Where(squirrel.Eq{"id": in.ID})

	return builder.PlaceholderFormat(squirrel.Dollar).ToSql()
}

func buildExamplesDelete(opts *option.Option) (string, []any, error) {
	builder := squirrel.Delete(exampleTable)

	builder = model2.AddDeleteFilter(builder, opts)

	return builder.PlaceholderFormat(squirrel.Dollar).ToSql()
}
