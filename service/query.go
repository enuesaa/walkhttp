package service

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type QueryService struct {}

func (srv *QueryService) Def() string {
	def := `
	type Query {
		hey: String!
	}
	`
	return def
}

func (srv *QueryService) QuerySchema() *graphql.Schema {
	return graphql.MustParseSchema(srv.Def(), &QueryResolver{})
}

func (srv *QueryService) Exec(query string, operationName string, variables map[string]interface{}) *graphql.Response {
	schema := srv.QuerySchema()
	return schema.Exec(context.Background(), query, operationName, variables)
}

type QueryResolver struct {}

func (_ *QueryResolver) Hey() string {
	return "Hey, world!"
}
