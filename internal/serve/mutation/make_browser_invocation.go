package mutation

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *MutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (*bool, error) {
	success := false

	invokeSrv := invoke.New(r.Repos)
	data := invokeSrv.Create(invocation.Method, invocation.URL)
	data.Request.Body = invocation.RequestBody
	data.Response.Status = invocation.Status
	data.Response.Body = invocation.ResponseBody

	if err := invokeSrv.Save(data); err != nil {
		return &success, err
	}
	success = true

	return &success, nil
}
