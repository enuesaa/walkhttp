package mutation

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/router/schema"
)

func (r *MutationResolver) MakeServerInvocation(ctx context.Context, invocation schema.ServerInvocationInput) (bool, error) {
	invokeSrv := invoke.New(r.Repos)
	entry := invocation.ToEntry()

	if err := invokeSrv.Invoke(&entry); err != nil {
		return false, err
	}
	if err := invokeSrv.Write(entry); err != nil {
		return false, err
	}
	return true, nil
}
