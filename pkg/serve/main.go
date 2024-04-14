package serve

import (
	"fmt"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/walkin/pkg/graph"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
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
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	
	// see https://github.com/99designs/gqlgen/issues/1664
	gqhandle := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	gqhandle.AddTransport(&transport.Websocket{})
	app.All("/graph", adaptor.HTTPHandlerFunc(gqhandle.ServeHTTP))
	app.Get("/graph/playground", adaptor.HTTPHandlerFunc(playground.Handler("graph", "/graph")))
	app.All("/*", ctlweb.Serve)

	return app.Listen(s.Addr())
}
