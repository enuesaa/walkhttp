package query

import (
	"context"
	"fmt"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *QueryResolver) GetInvocation(ctx context.Context, id string) (*schema.Invocation, error) {
	invokeSrv := invoke.New(r.Repos)
	entry, err := invokeSrv.Get(id)
	if err != nil {
		return nil, err
	}
	
	invocation := schema.Invocation{
		ID: entry.Id,
		Status: entry.Response.Status,
		Method: entry.Request.Method,
		URL: entry.Request.Url,
		RequestHeaders: make([]*schema.Header, 0),
		ResponseHeaders: make([]*schema.Header, 0),
		RequestBody: entry.Request.Body,
		ResponseBody: entry.Response.Body,
		Created: fmt.Sprint(entry.Request.Started),
	}
	return &invocation, nil
}
