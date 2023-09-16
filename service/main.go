package service

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/enuesaa/walkin/resolver"
)

func ExecQuery(query string, operationName string, variables map[string]interface{}) (*graphql.Response, error) {
	schema, err := graphql.ParseSchema(resolver.Schema, &resolver.Resolver{})
	if err != nil {
		return nil, err
	}
	return schema.Exec(context.Background(), query, operationName, variables), nil
}
