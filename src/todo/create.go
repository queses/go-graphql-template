package todo

import (
	"context"
	"errors"
)

type Create struct {
	repo TodoRepo
}

func NewCreate(repo TodoRepo) *Create {
	return &Create{repo: repo}
}

func (uc *Create) Run(ctx context.Context, text string) (*TodoEntity, error) {
	if text == "" || len(text) > 2048 {
		return nil, errors.New("Bad text format")
	}

	return uc.repo.CreateTodo(ctx, text)
}
