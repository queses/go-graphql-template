//go:generate go run github.com/99designs/gqlgen generate

package resolvers

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/queses/go-graphql-template/src/graph"
	"github.com/queses/go-graphql-template/src/lib"
	"github.com/rs/zerolog/log"
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

func transformError(ctx context.Context, err error) error {
	internalErr, isInternalErr := err.(internalError)
	if !isInternalErr {
		return err
	}

	errMessage := fmt.Sprintf("%s%+v", internalErr.Error(), internalErr.StackTrace()[0:10])
	log.Ctx(ctx).Error().Msg(errMessage)
	return errors.New("internal error")

}
