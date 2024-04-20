package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/enuesaa/walkin/pkg/graph"
)

func Serve(repos repository.Repos, port int) error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// see https://github.com/99designs/gqlgen/issues/1664
	gqhandle := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	gqhandle.AddTransport(&transport.Websocket{})
	app.Any("/graph", echo.WrapHandler(gqhandle))
	app.GET("/graph/playground", echo.WrapHandler(playground.Handler("graph", "/graph")))
	app.GET("/aa", func(c echo.Context) error {
		invocation := invoke.NewInvocation("GET", "https://example.com")
		return invoke.Invoke(&invocation)
	})

	return app.Start(fmt.Sprintf(":%d", port))
}
