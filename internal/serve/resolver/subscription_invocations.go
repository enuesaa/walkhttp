package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *subscriptionResolver) SubscribeInvocations(ctx context.Context) (<-chan []*schema.Invocation, error) {
	invokeSrv := invoke.New(r.Repos)

	ch := make(chan []*schema.Invocation)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)

			invocations := make([]*schema.Invocation, 0)
			ws, err := invokeSrv.Read()
			if err != nil {
				continue
			}
			
			for _, entry := range ws.Entries {
				invocations = append(invocations, &schema.Invocation{
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
			select {
			case <-ctx.Done():
				return
			case ch <- invocations:
			}
		}
	}()

	return ch, nil
}
