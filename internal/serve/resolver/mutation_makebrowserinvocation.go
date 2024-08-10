package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *mutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (*bool, error) {
	invokeSrv := invoke.New(r.Repos)
	data := invokeSrv.Create(invocation.Method, invocation.URL)

	data.Request.Body = invocation.RequestBody
	data.Response.Status = invocation.Status
	data.Response.Body = invocation.ResponseBody

	if err := invokeSrv.Save(data); err != nil {
		return nil, err
	}
	return nil, nil
}
