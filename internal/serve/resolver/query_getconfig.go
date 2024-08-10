package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/internal/serve/schema"
)

func (r *queryResolver) GetConfig(ctx context.Context) (*schema.Config, error) {
	config := schema.Config{
		BaseURL: r.BaseUrl,
	}
	return &config, nil
}