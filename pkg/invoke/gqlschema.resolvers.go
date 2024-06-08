package invoke

// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"time"
)

func (r *queryResolver) Invocations(ctx context.Context) ([]*Invocation, error) {
	list := make([]*Invocation, 0)
	ids := r.repos.DB.List()
	for _, id := range ids {
		data, err := r.repos.DB.Get(id)
		if err != nil {
			return list, err
		}
		invocation := data.(Invocation)
		list = append(list, &invocation)
	}
	return list, nil
}

func (r *queryResolver) Invocation(ctx context.Context, id string) (*Invocation, error) {
	data, err := r.repos.DB.Get(id)
	if err != nil {
		return nil, err
	}
	invocation := data.(Invocation)
	return &invocation, nil
}

func (r *subscriptionResolver) Invocations(ctx context.Context) (<-chan []*Invocation, error) {
	ch := make(chan []*Invocation)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)

			invocations := make([]*Invocation, 0)
			ids := r.repos.DB.List()
			for _, id := range ids {
				data, err := r.repos.DB.Get(id)
				if err != nil {
					continue
				}
				invocation := data.(Invocation)
				invocations = append(invocations, &invocation)
			}
			select {
			case <-ctx.Done():
				return
			case ch <- invocations:
			}
		}
	}()

	return ch, nil
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *mutationResolver) MakeServerInvocation(ctx context.Context, invocation ServerInvocationInput) (*bool, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *mutationResolver) MakeBrowserInvocation(ctx context.Context, invocation BrowserInvocationInput) (*bool, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}
