package service

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

// see https://github.com/graph-gophers/graphql-go/blob/master/example/starwars/starwars.go#L146
type QueryResolver struct {}

func (_ *QueryResolver) Hey() string {
	return "Hey, world!"
}
func (r *QueryResolver) Fileinfo(args struct{ Name string }) *FileinfoResolver {
	if args.Name == "aa" {
		return &FileinfoResolver{}
	}
	return nil
}

type FileinfoResolver struct {}
func (_ *FileinfoResolver) Name() string {
	return "aa"
}
func (_ *FileinfoResolver) Description() string {
	return "aa-description"
}

func ExecQuery(query string, operationName string, variables map[string]interface{}) *graphql.Response {
	queryDef := `
	schema {
		query: Query
	}
	type Query {
		fileinfo(name: String!): Fileinfo
	}
	type Fileinfo {
		name: String!
		description: String!
	}
	`
	schema := graphql.MustParseSchema(queryDef, &QueryResolver{})
	return schema.Exec(context.Background(), query, operationName, variables)
}
