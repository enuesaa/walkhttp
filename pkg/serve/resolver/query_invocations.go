package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

func (r *queryResolver) Invocations(ctx context.Context) ([]*schema.Invocation, error) {
	list := make([]*schema.Invocation, 0)
	ids := r.Repos.DB.List()
	for _, id := range ids {
		data, err := r.Repos.DB.Get(id)
		if err != nil {
			return list, err
		}
		invocation := data.(schema.Invocation)
		list = append(list, &invocation)
	}
	return list, nil
}
