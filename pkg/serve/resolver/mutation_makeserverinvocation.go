package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

func (r *mutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (*bool, error) {
	invokeSrv := invoke.New(r.Repos)
	data := invokeSrv.Create(invocation.Method, invocation.URL)
	if err := invokeSrv.Invoke(&data); err != nil {
		return nil, err
	}
	if err := r.Repos.DB.Save(data.ID, data); err != nil {
		return nil, err
	}
	return nil, nil
}