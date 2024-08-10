package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *mutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (*bool, error) {
	invokeSrv := invoke.New(r.Repos)
	data := invokeSrv.Create(invocation.Method, invocation.URL)
	if err := invokeSrv.Invoke(&data); err != nil {
		return nil, err
	}
	if err := invokeSrv.Save(data); err != nil {
		return nil, err
	}
	return nil, nil
}