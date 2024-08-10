package resolver

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
	// "github.com/google/uuid"
)

func (r *mutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (*bool, error) {
	// id := uuid.NewString()

	data := workspace.Entry{
		Request: workspace.EntryRequest{
			Method: invocation.Method,
			Url: invocation.URL,
			Headers: map[string][]string{},
			Body: invocation.RequestBody,
			Started: time.Now().Unix(),
		},
		Response: workspace.EntryResponse{
			Status: invocation.Status,
			Body: invocation.ResponseBody,
		},
	}
	invokeSrv := invoke.New(r.Repos)
	if err := invokeSrv.Save(data); err != nil {
		return nil, err
	}
	return nil, nil
}
