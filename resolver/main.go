package resolver

// see https://github.com/graph-gophers/graphql-go/blob/master/example/starwars/starwars.go#L146
var Schema = `
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