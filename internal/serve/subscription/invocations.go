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
			entries, err := invokeSrv.List()
			if err != nil {
				continue
			}

			list := make([]*schema.Invocation, 0)
			for _, entry := range entries {
				invocation := schema.NewInvocationFromEntry(entry)
				list = append(list, &invocation)
			}

			select {
			case <-ctx.Done():
				return
			case ch <- list:
			}
		}
	}()

	return ch, nil
}
