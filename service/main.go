package service

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/enuesaa/walkin/resolver"
)

func ExecQuery(query string, operationName string, variables map[string]interface{}) *graphql.Response {
	schema := graphql.MustParseSchema(resolver.Schema, &resolver.QueryResolver{})
	return schema.Exec(context.Background(), query, operationName, variables)
}
