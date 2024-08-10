package query

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *QueryResolver) GetConfig(ctx context.Context) (*schema.Config, error) {
	invokeSrv := invoke.New(r.Repos)
	
	config := schema.Config{
		BaseURL: invokeSrv.GetBaseUrl(),
	}
	return &config, nil
}
