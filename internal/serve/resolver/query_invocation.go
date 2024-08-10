package resolver

import (
	"context"
	"fmt"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *queryResolver) GetInvocation(ctx context.Context, id string) (*schema.Invocation, error) {
	invokeSrv := invoke.New(r.Repos)

	ws, err := invokeSrv.Read()
	if err != nil {
		return &schema.Invocation{}, err
	}

	for _, entry := range ws.Entries {
		if entry.Id == id {
			data := schema.Invocation{
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
			return &data, nil
		}
	}
	
	return &schema.Invocation{}, fmt.Errorf("not found")
}
