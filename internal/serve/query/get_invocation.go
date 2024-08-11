package query

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *QueryResolver) GetInvocation(ctx context.Context, id string) (*schema.Invocation, error) {
	invokeSrv := invoke.New(r.Repos)
	entry, err := invokeSrv.Get(id)
	if err != nil {
		return nil, err
	}

	invocation := schema.NewInvocationFromEntry(entry)

	return &invocation, nil
}
