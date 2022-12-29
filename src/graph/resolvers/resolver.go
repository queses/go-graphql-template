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

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func onError(err error) {
	var log string
	withStack, ok := err.(stackTracer)
	if ok {
		log = fmt.Sprintf("%s%+v", err.Error(), withStack.StackTrace()[0:10])
	} else {
		log = err.Error()
	}

	fmt.Println(log)
}
