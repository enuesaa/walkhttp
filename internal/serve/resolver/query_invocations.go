package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *queryResolver) ListInvocations(ctx context.Context) ([]*schema.Invocation, error) {
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
