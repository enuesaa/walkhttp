package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/enuesaa/walkin/pkg/graph/model"
)

// Invocatoins is the resolver for the invocatoins field.
func (r *queryResolver) Invocatoins(ctx context.Context) ([]*model.Invocation, error) {
	list := make([]*model.Invocation, 0)
	list = append(list, &model.Invocation{
		ID:     "a",
		Status: 200,
		Method: "GET",
		URL:    "https://example.com",
	})
	return list, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
