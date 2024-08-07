package serve

import (
	"net/http"
	"time"

	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/serve/gql"
	"github.com/enuesaa/walkhttp/internal/serve/resolver"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func ServeGQ(repos repository.Repos, baseUrl string, port int) echo.HandlerFunc {
	// see https://github.com/99designs/gqlgen/issues/1664
	// see https://github.com/99designs/gqlgen/issues/2826
	gqhandle := handler.New(gql.NewExecutableSchema(gql.Config{
		Resolvers: &resolver.Resolver{
			Repos:   repos,
			BaseUrl: baseUrl,
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

	return echo.WrapHandler(gqhandle)
}
