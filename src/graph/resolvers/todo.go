package resolvers

import (
	"context"
	"fmt"
	"github.com/queses/go-graphql-template/src/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var items []*model.Todo
	dummyItem := model.Todo{
		ID:   "Item1",
		Text: "Do something",
		Done: false,
		User: nil,
	}

	items = append(items, &dummyItem)
	return items, nil
}
