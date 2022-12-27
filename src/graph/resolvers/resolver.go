//go:generate go run github.com/99designs/gqlgen generate

package resolvers

import (
	"github.com/jmoiron/sqlx"
	"github.com/queses/go-graphql-template/src/graph"
)

type Resolver struct {
	Db *sqlx.DB
}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
