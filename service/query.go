package service

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

// todo define queries
type QueryService struct {}

func (srv *QueryService) QuerySchema() *graphql.Schema {
	return graphql.MustParseSchema(QueryDef, &QueryResolver{})
}

func (srv *QueryService) Exec(query string, operationName string, variables map[string]interface{}) *graphql.Response {
	schema := srv.QuerySchema()
	return schema.Exec(context.Background(), query, operationName, variables)
}

// 生成できないかな
var QueryDef = `
type Query {
	hey: String!
}
`

type QueryResolver struct {}

func (_ *QueryResolver) Hey() string {
	return "Hey, world!"
}
