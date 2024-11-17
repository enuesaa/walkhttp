package mutation

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/routegql/schema"
)

func (r *MutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (bool, error) {
	invokeSrv := invoke.New(r.Repos)

	entry := invokeSrv.NewEntry(invocation.Method, "")

	// req
	entry.Request.Body = invocation.RequestBody
	entry.Request.Url = invocation.URL
	entry.Request.Started = time.Unix(0, 0).Unix()

	for _, h := range invocation.RequestHeaders {
		entry.Request.Headers[h.Name] = h.Value
	}

	// invoke
	if err := invokeSrv.Invoke(&entry); err != nil {
		return false, err
	}

	// save
	if err := invokeSrv.Save(entry); err != nil {
		return false, err
	}

	return true, nil
}
