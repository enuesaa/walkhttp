package serve

import (
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/serve/gql"
)

type Resolver struct {
	repos   repository.Repos
	baseUrl string
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Subscription() gql.SubscriptionResolver {
	return &subscriptionResolver{r}
}
