package lib

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/queses/go-graphql-template/src/todo"
	"github.com/queses/go-graphql-template/src/todo/postgres"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

type ServiceFactory struct {
	pg       *sqlx.DB
	TodoRepo todo.TodoRepo
}

func NewServiceFactory() *ServiceFactory {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

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

func (f *ServiceFactory) GetLogCtx(ctx context.Context, operationId string) context.Context {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	return logger.With().Str("operationId", operationId).Logger().WithContext(ctx)
}
