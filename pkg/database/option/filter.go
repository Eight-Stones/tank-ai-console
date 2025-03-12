package option

type Filter interface {
	Query() (string, []any)
}
