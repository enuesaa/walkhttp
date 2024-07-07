package invoke

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/pkg/invoke/gql"
	"github.com/enuesaa/walkhttp/pkg/invoke/schema"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/google/uuid"
)

type Resolver struct {
	repos repository.Repos
	baseUrl string
}

type queryResolver struct {
	*Resolver
}

type mutationResolver struct {
	*Resolver
}

type subscriptionResolver struct {
	*Resolver
}

//TODO refactor
func (r *queryResolver) AppConfig(ctx context.Context) (*schema.AppConfig, error) {
	config := schema.AppConfig{
		BaseURL: r.baseUrl,
	}
	return &config, nil
}

func (r *queryResolver) Invocations(ctx context.Context) ([]*schema.Invocation, error) {
	list := make([]*schema.Invocation, 0)
	ids := r.repos.DB.List()
	for _, id := range ids {
		data, err := r.repos.DB.Get(id)
		if err != nil {
			return list, err
		}
		invocation := data.(schema.Invocation)
		list = append(list, &invocation)
	}
	return list, nil
}

func (r *queryResolver) Invocation(ctx context.Context, id string) (*schema.Invocation, error) {
	data, err := r.repos.DB.Get(id)
	if err != nil {
		return nil, err
	}
	invocation := data.(schema.Invocation)
	return &invocation, nil
}

func (r *subscriptionResolver) Invocations(ctx context.Context) (<-chan []*schema.Invocation, error) {
	ch := make(chan []*schema.Invocation)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)

			invocations := make([]*schema.Invocation, 0)
			ids := r.repos.DB.List()
			for _, id := range ids {
				data, err := r.repos.DB.Get(id)
				if err != nil {
					continue
				}
				invocation := data.(schema.Invocation)
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

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *mutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (*bool, error) {
	invokeSrv := NewInvokeSrv(r.repos)
	data := invokeSrv.Create(invocation.Method, invocation.URL)
	if err := invokeSrv.Invoke(&data); err != nil {
		return nil, err
	}
	if err := r.repos.DB.Save(data.ID, data); err != nil {
		return nil, err
	}
	return nil, nil
}
func (r *mutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (*bool, error) {
	id := uuid.NewString()
	data := schema.Invocation{
		ID:              id,
		Status:          invocation.Status,
		Method:          invocation.Method,
		URL:             invocation.URL,
		RequestHeaders:  make([]*schema.Header, 0),
		ResponseHeaders: make([]*schema.Header, 0),
		RequestBody:     invocation.RequestBody,
		ResponseBody:    invocation.ResponseBody,
		Created:         time.Now().Format(time.RFC3339),
	}
	if err := r.repos.DB.Save(id, data); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Subscription() gql.SubscriptionResolver {
	return &subscriptionResolver{r}
}
