package lib

import (
	"github.com/jmoiron/sqlx"
	"github.com/queses/go-graphql-template/src/todo"
	"github.com/queses/go-graphql-template/src/todo/postgres"
)

type ServiceFactory struct {
	pg       *sqlx.DB
	TodoRepo todo.TodoRepo
}

func NewServiceFactory() *ServiceFactory {
	pg, err := sqlx.Open(
		"postgres",
		"postgres://postgres:pass@127.0.0.1:54331/postgres?sslmode=disable",
	)

	if err != nil {
		panic(err)
	}

	factory := &ServiceFactory{
		pg:       pg,
		TodoRepo: postgres.NewTodoRepo(pg),
	}

	return factory
}

func (f *ServiceFactory) Close() {
	err := f.pg.Close()
	if err != nil {
		panic(err)
	}
}
