package model

type Example struct {
	ID   int64  `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
	Meta string `db:"meta"`
}
