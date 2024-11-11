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
	// request
	entry.Request.Body = invocation.RequestBody
	// TODO url should starts with WALKHTTP_URL_BASE
	entry.Request.Url = invocation.URL

	started, err := time.Parse(time.RFC3339, invocation.Started)
	if err != nil {
		started = time.Unix(0, 0)
	}
	entry.Request.Started = started.Unix()

	for _, h := range invocation.RequestHeaders {
		if _, ok := entry.Request.Headers[h.Name]; !ok {
			entry.Request.Headers[h.Name] = []string{}
		}
		entry.Request.Headers[h.Name] = append(entry.Request.Headers[h.Name], h.Value)
	}

	// response
	entry.Response.Status = invocation.Status
	entry.Response.Body = invocation.ResponseBody

	received, err := time.Parse(time.RFC3339, invocation.Received)
	if err != nil {
		received = time.Unix(0, 0)
	}
	entry.Response.Received = received.Unix()

	for _, h := range invocation.ResponseHeaders {
		if _, ok := entry.Response.Headers[h.Name]; !ok {
			entry.Response.Headers[h.Name] = []string{}
		}
		entry.Response.Headers[h.Name] = append(entry.Response.Headers[h.Name], h.Value)
	}

	if err := invokeSrv.Save(entry); err != nil {
		return false, err
	}
	return true, nil
}
