package mutation

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *MutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (bool, error) {
	invokeSrv := invoke.New(r.Repos)
	entry := invocation.ToEntry()

	if err := invokeSrv.Write(entry); err != nil {
		return false, err
	}
	return true, nil
}
