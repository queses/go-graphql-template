package model

type TodoRow struct {
	Id   string `db:"id"`
	Text string `db:"text"`
	Done bool   `db:"done"`
}
