package query

import (
	"context"
	"fmt"

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
		list = append(list, &schema.Invocation{
			ID: entry.Id,
			Status: entry.Response.Status,
			Method: entry.Request.Method,
			URL: entry.Request.Url,
			RequestHeaders: make([]*schema.Header, 0),
			ResponseHeaders: make([]*schema.Header, 0),
			RequestBody: entry.Request.Body,
			ResponseBody: entry.Response.Body,
			Created: fmt.Sprint(entry.Request.Started),
		})
	}
	
	return list, nil
}
