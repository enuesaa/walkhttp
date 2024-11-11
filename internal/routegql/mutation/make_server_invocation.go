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
	// request
	entry.Request.Body = invocation.RequestBody
	// TODO url should starts with WALKHTTP_URL_BASE
	entry.Request.Url = invocation.URL
	entry.Request.Started = time.Unix(0, 0).Unix()

	for _, h := range invocation.RequestHeaders {
		if _, ok := entry.Request.Headers[h.Name]; !ok {
			entry.Request.Headers[h.Name] = []string{}
		}
		entry.Request.Headers[h.Name] = append(entry.Request.Headers[h.Name], h.Value)
	}

	if err := invokeSrv.Invoke(&entry); err != nil {
		return false, err
	}
	if err := invokeSrv.Save(entry); err != nil {
		return false, err
	}
	return true, nil
}
