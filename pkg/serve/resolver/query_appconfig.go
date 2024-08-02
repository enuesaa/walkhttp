package resolver

import (
	"context"

	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

// TODO refactor
func (r *queryResolver) AppConfig(ctx context.Context) (*schema.AppConfig, error) {
	config := schema.AppConfig{
		BaseURL: r.BaseUrl,
	}
	return &config, nil
}