package serve

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/serve/gql"
	"github.com/enuesaa/walkhttp/internal/serve/mutation"
	"github.com/enuesaa/walkhttp/internal/serve/query"
	"github.com/enuesaa/walkhttp/internal/serve/subscription"
)

type Resolver struct {
	Repos   repository.Repos
	BaseUrl string
}

func (r *Resolver) Query() gql.QueryResolver {
	resolver := query.QueryResolver{
		Repos: r.Repos,
	}
	return &resolver
}

func (r *Resolver) Mutation() gql.MutationResolver {
	resolver := mutation.MutationResolver{
		Repos: r.Repos,
	}
	return &resolver
}

func (r *Resolver) Subscription() gql.SubscriptionResolver {
	resolver := subscription.SubscriptionResolver{
		Repos: r.Repos,
	}
	return &resolver
}
