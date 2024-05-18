package invoke

import (
	"fmt"
	"net/http"
	"time"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

func ServeGraphql(repos repository.Repos, port int) error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// see https://github.com/99designs/gqlgen/issues/1664
	// see https://github.com/99designs/gqlgen/issues/2826
	gqhandle := handler.New(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			repos: repos,
		},
	}))
	gqhandle.AddTransport(transport.Options{})
	gqhandle.AddTransport(transport.GET{})
	gqhandle.AddTransport(transport.POST{})
	gqhandle.AddTransport(transport.MultipartForm{})
	gqhandle.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	gqhandle.Use(extension.Introspection{})
	app.Any("/graph", echo.WrapHandler(gqhandle))
	app.GET("/graph/playground", echo.WrapHandler(playground.Handler("graph", "/graph")))
	app.Any("/*", ctlweb.Serve)
	app.Any("/", ctlweb.Serve)

	return app.Start(fmt.Sprintf(":%d", port))
}
