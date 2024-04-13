package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/walkin/pkg/graph"
)

func CreateGraphCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "poc graphql server",
		RunE: func(cmd *cobra.Command, args []string) error {
			schema := graph.NewExecutableSchema(graph.Config{
				Resolvers: &graph.Resolver{},
			})
			srv := handler.NewDefaultServer(schema)
	
			http.Handle("/", playground.Handler("GraphQL playground", "/query"))
			http.Handle("/query", srv)
			if err := http.ListenAndServe(":3002", nil); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}

			return nil
		},
	}

	return cmd
}
