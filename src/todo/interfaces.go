package todo

import "context"

type TodoRepo interface {
	CreateTodo(ctx context.Context, text string) (*TodoEntity, error)
	ListTodos(ctx context.Context) ([]TodoEntity, error)
}
