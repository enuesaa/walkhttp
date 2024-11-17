package mutation

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/routegql/schema"
)

func (r *MutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (bool, error) {
	invokeSrv := invoke.New(r.Repos)

	entry := invokeSrv.NewEntry(invocation.Method, "")

	// req
	entry.Request.Body = invocation.RequestBody
	entry.Request.Url = invocation.URL

	for _, h := range invocation.RequestHeaders {
		entry.Request.Headers[h.Name] = h.Value
	}

	started, err := time.Parse(time.RFC3339, invocation.Started)
	if err != nil {
		started = time.Unix(0, 0)
	}
	entry.Request.Started = started.Unix()

	// res
	entry.Response.Status = invocation.Status
	entry.Response.Body = invocation.ResponseBody

	for _, h := range invocation.ResponseHeaders {
		entry.Response.Headers[h.Name] = h.Value
	}

	received, err := time.Parse(time.RFC3339, invocation.Received)
	if err != nil {
		received = time.Unix(0, 0)
	}
	entry.Response.Received = received.Unix()


	// save
	if err := invokeSrv.Save(entry); err != nil {
		return false, err
	}

	return true, nil
}
