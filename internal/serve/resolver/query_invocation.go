package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *queryResolver) GetInvocation(ctx context.Context, id string) (*schema.Invocation, error) {
	return &schema.Invocation{}, nil
	
	// data, err := r.Repos.DB.Get(id)
	// if err != nil {
	// 	return nil, err
	// }
	// invocation := data.(schema.Invocation)
	// return &invocation, nil
}
