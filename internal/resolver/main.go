package resolver

// see https://github.com/graph-gophers/graphql-go/blob/master/example/starwars/starwars.go#L146
var Schema = `
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

func (r *Resolver) Fileinfo(args struct{ Name string }) *FileinfoResolver {
	// run os
	if args.Name == "aa" {
		return &FileinfoResolver{ name: "aa", description: "aa-description" }
	}
	return &FileinfoResolver{ name: "aa", description: "aa-description" }
}

func (r *Resolver) CreateFile(args struct{ Name string }) *FileinfoResolver {
	return &FileinfoResolver{ name: "aa", description: "aa-description" }
}
