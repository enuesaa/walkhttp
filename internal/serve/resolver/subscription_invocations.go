package resolver

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *subscriptionResolver) SubscribeInvocations(ctx context.Context) (<-chan []*schema.Invocation, error) {
	ch := make(chan []*schema.Invocation)

	go func() {
		defer close(ch)
		for {
			time.Sleep(1 * time.Second)

			invocations := make([]*schema.Invocation, 0)
			ids := r.Repos.DB.List()
			for _, id := range ids {
				data, err := r.Repos.DB.Get(id)
				if err != nil {
					continue
				}
				invocation := data.(invoke.Invocation)
				invocations = append(invocations, &schema.Invocation{
					ID: invocation.ID,
					Status: invocation.Status,
					Method: invocation.Method,
					URL: invocation.URL,
					RequestHeaders: make([]*schema.Header, 0),
					ResponseHeaders: make([]*schema.Header, 0),
					RequestBody: invocation.RequestBody,
					ResponseBody: invocation.ResponseBody,
					Created: invocation.Created,
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
