package todo

type TodoEntity struct {
	Id   string `db:"id"`
	Text string `db:"text"`
	Done bool   `db:"done"`
}
