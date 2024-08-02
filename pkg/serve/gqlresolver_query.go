package serve

import (
	"context"

	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

type queryResolver struct {
	*Resolver
}

// TODO refactor
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
