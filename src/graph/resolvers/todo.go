package resolvers

import (
	"context"
	model2 "github.com/queses/go-graphql-template/src/db/model"
	"github.com/queses/go-graphql-template/src/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	var rows []model2.TodoRow
	err := r.Db.Select(&rows, "INSERT INTO todo (text) VALUES ($1) RETURNING id, text, done", input.Text)
	if err != nil {
		return nil, err
	}

	item := &model.Todo{
		ID:   rows[0].Id,
		Text: rows[0].Text,
		Done: rows[0].Done,
	}

	return item, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var rows []model2.TodoRow
	err := r.Db.Select(&rows, "SELECT id, text, done FROM todo ORDER BY createdAt ASC")
	if err != nil {
		return nil, err
	}

	items := make([]*model.Todo, len(rows))
	for i, row := range rows {
		items[i] = &model.Todo{
			ID:   row.Id,
			Text: row.Text,
			Done: row.Done,
		}
	}

	return items, nil
}
