package subscription

import (
	"context"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *SubscriptionResolver) SubscribeInvocations(ctx context.Context) (<-chan []*schema.Invocation, error) {
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
				invocation := schema.NewInvocationFromEntry(entry)
				invocations = append(invocations, &invocation)
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
