package serve

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/serve/schema"
	"github.com/google/uuid"
)

type mutationResolver struct {
	*Resolver
}

func (r *mutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (*bool, error) {
	invokeSrv := invoke.New(r.repos)
	data := invokeSrv.Create(invocation.Method, invocation.URL)
	if err := invokeSrv.Invoke(&data); err != nil {
		return nil, err
	}
	if err := r.repos.DB.Save(data.ID, data); err != nil {
		return nil, err
	}
	return nil, nil
}
func (r *mutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (*bool, error) {
	id := uuid.NewString()
	data := schema.Invocation{
		ID:              id,
		Status:          invocation.Status,
		Method:          invocation.Method,
		URL:             invocation.URL,
		RequestHeaders:  make([]*schema.Header, 0),
		ResponseHeaders: make([]*schema.Header, 0),
		RequestBody:     invocation.RequestBody,
		ResponseBody:    invocation.ResponseBody,
		Created:         time.Now().Format(time.RFC3339),
	}
	if err := r.repos.DB.Save(id, data); err != nil {
		return nil, err
	}
	return nil, nil
}
