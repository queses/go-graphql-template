//go:generate go run github.com/99designs/gqlgen generate

package resolvers

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/queses/go-graphql-template/src/graph"
	"github.com/queses/go-graphql-template/src/lib"
)

type Resolver struct {
	Factory *lib.ServiceFactory
}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

type internalError interface {
	error
	StackTrace() errors.StackTrace
}

func transformError(err error) error {
	var log string
	internalErr, ok := err.(internalError)
	if !ok {
		return err
	}

	log = fmt.Sprintf("%s%+v", internalErr.Error(), internalErr.StackTrace()[0:10])
	fmt.Println(log)

	return errors.New("internal error")
}
