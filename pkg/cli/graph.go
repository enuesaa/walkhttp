package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/walkin/graph"
)

func CreateGraphCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "graph",
		Short: "poc graphql server",
		RunE: func(cmd *cobra.Command, args []string) error {
			srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	
			http.Handle("/", playground.Handler("GraphQL playground", "/query"))
			http.Handle("/query", srv)
			log.Fatal(http.ListenAndServe(":3002", nil))

			return nil
		},
	}

	return cmd
}
