package routegql

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/routegql/mutation"
	"github.com/enuesaa/walkhttp/internal/routegql/query"
	"github.com/enuesaa/walkhttp/internal/routegql/subscription"
)

type Resolver struct {
	Repos repository.Repos
}

func (r *Resolver) Query() QueryResolver {
	resolver := query.QueryResolver{
		Repos: r.Repos,
	}
	return &resolver
}

func (r *Resolver) Mutation() MutationResolver {
	resolver := mutation.MutationResolver{
		Repos: r.Repos,
	}
	return &resolver
}

func (r *Resolver) Subscription() SubscriptionResolver {
	resolver := subscription.SubscriptionResolver{
		Repos: r.Repos,
	}
	return &resolver
}
