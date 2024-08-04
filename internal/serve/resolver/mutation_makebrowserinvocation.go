package resolver

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/internal/serve/schema"
	"github.com/google/uuid"
)

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
	if err := r.Repos.DB.Save(id, data); err != nil {
		return nil, err
	}
	return nil, nil
}
