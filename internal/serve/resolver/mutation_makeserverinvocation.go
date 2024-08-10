package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *mutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (*bool, error) {
	success := false

	invokeSrv := invoke.New(r.Repos)
	data := invokeSrv.Create(invocation.Method, invocation.URL)
	data.Request.Body = invocation.RequestBody

	if err := invokeSrv.Invoke(&data); err != nil {
		return &success, err
	}
	if err := invokeSrv.Save(data); err != nil {
		return &success, err
	}
	success = true

	return &success, nil
}
