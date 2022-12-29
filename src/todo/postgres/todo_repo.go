package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/queses/go-graphql-template/src/todo"
)

type TodoRepo struct {
	db *sqlx.DB
}

func NewTodoRepo(db *sqlx.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) CreateTodo(ctx context.Context, text string) (*todo.TodoEntity, error) {
	var rows []todo.TodoEntity
	err := r.db.Select(&rows, "INSERT INTO todo (text) VALUES ($1) RETURNING id, text, done", text)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &rows[0], nil
}

func (r *TodoRepo) ListTodos(ctx context.Context) ([]todo.TodoEntity, error) {
	var rows []todo.TodoEntity
	err := r.db.Select(&rows, "SELECT id, text, done FROM todo ORDER BY createdAt ASC")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return rows, nil
}
