package serve

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func ServeGQPlayground() echo.HandlerFunc {
	return echo.WrapHandler(playground.Handler("graph", "/graph"))
}
