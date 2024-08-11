package mutation

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *MutationResolver) MakeBrowserInvocation(ctx context.Context, invocation schema.BrowserInvocationInput) (*bool, error) {
	success := false

	invokeSrv := invoke.New(r.Repos)
	entry := invocation.ToEntry()

	if err := invokeSrv.Save(entry); err != nil {
		return &success, err
	}
	success = true

	return &success, nil
}
