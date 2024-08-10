package resolver

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *mutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (*bool, error) {
	invokeSrv := invoke.New(r.Repos)

	data := workspace.Entry{
		Request: workspace.EntryRequest{
			Method: invocation.Method,
			Url: invocation.URL,
			Headers: map[string][]string{},
			Body: invocation.RequestBody,
			Started: time.Now().Unix(),
		},
		Response: workspace.EntryResponse{},
	}
	if err := invokeSrv.Invoke(&data); err != nil {
		return nil, err
	}
	if err := invokeSrv.Save(data); err != nil {
		return nil, err
	}
	return nil, nil
}