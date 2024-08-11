package serve

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func (ctl *ServeCtl) handleGqlPlayground() echo.HandlerFunc {
	return echo.WrapHandler(playground.Handler("graph", "/graph"))
}
