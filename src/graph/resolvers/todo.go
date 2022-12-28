package resolvers

import (
	"context"
	"github.com/queses/go-graphql-template/src/graph/model"
	"github.com/queses/go-graphql-template/src/todo"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	entity, err := todo.NewCreate(r.Factory.TodoRepo).Run(ctx, input.Text)
	if err != nil {
		return nil, err
	}

	item := &model.Todo{ID: entity.Id, Text: entity.Text, Done: entity.Done}

	return item, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	entities, err := todo.NewList(r.Factory.TodoRepo).Run(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*model.Todo, len(entities))
	for i, entity := range entities {
		items[i] = &model.Todo{ID: entity.Id, Text: entity.Text, Done: entity.Done}
	}

	return items, nil
}
