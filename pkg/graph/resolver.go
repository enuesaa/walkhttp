package graph

import (
	"context"

	"github.com/enuesaa/walkin/pkg/graph/model"
)

type Resolver struct{}
func (r *Resolver) Invocatoins(ctx context.Context) ([]*model.Invocation, error) {
	list := make([]*model.Invocation, 0)
	list = append(list, &model.Invocation{
		ID: "a",
		Status: 200,
		Method: "GET",
		URL: "https://example.com",
	})
	return list, nil
}
