package resolver

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/serve/gql"
)

type Resolver struct {
	Repos   repository.Repos
	BaseUrl string
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
