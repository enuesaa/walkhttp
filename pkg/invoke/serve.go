package invoke

import (
	"fmt"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(repos repository.Repos, port int) error {
	app := echo.New()
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	app.Any("/graph", ServeGQ(repos, port))
	app.GET("/graph/playground", ServeGQPlayground())
	app.Any("/*", ctlweb.Serve)
	app.Any("/", ctlweb.Serve)

	return app.Start(fmt.Sprintf(":%d", port))
}
