package resolver

// see https://github.com/graph-gophers/graphql-go/blob/master/example/starwars/starwars.go#L146
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	type Query {
		fileinfo(name: String!): Fileinfo
	}
	type Mutation {
		createFile(name: String!): Fileinfo
	}

	type Fileinfo {
		name: String!
		description: String!
	}
`

type Resolver struct {}

func (_ *Resolver) Query() *QueryResolver {
	return &QueryResolver{}
}

func (_ *Resolver) Mutation() *MutationResolver {
	return &MutationResolver{}
}