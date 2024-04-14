package serve

import (
	"fmt"
	"net/http"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/walkin/pkg/graph"
)

func New(repos repository.Repos) Servectl {
	return Servectl{
		repos:   repos,
		wsconns: NewWsConns(),
		Port:    3000,
	}
}

type Servectl struct {
	repos   repository.Repos
	wsconns WsConns
	Port    int
}

func (s *Servectl) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *Servectl) Listen() error {
	// see https://github.com/99designs/gqlgen/issues/1664
	gqhandle := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	gqhandle.AddTransport(&transport.Websocket{})
	http.Handle("/graph/playground", playground.Handler("graph", "/graph"))
	http.Handle("/graph", gqhandle)

	return http.ListenAndServe(s.Addr(), nil)
}
