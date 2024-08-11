package query

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *QueryResolver) ListInvocations(ctx context.Context) ([]*schema.Invocation, error) {
	list := make([]*schema.Invocation, 0)
	invokeSrv := invoke.New(r.Repos)

	entries, err := invokeSrv.List()
	if err != nil {
		return list, err
	}

	for _, entry := range entries {
		invocation := schema.NewInvocationFromEntry(entry)
		list = append(list, &invocation)
	}
	
	return list, nil
}
