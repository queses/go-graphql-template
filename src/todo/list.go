package todo

import (
	"context"
)

type List struct {
	repo TodoRepo
}

func NewList(repo TodoRepo) *List {
	return &List{repo: repo}
}

func (uc *List) Run(ctx context.Context) ([]TodoEntity, error) {
	return uc.repo.ListTodos(ctx)
}
