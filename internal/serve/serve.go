package serve

import (
	"fmt"

	"github.com/enuesaa/walkhttp/ctlweb"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(repos repository.Repos, port int) error {
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	app.Any("/graph", ServeGQ(repos, port))
	app.GET("/graph/playground", ServeGQPlayground())
	app.Any("/*", ctlweb.Serve)
	app.Any("/", ctlweb.Serve)

	app.HideBanner = true
	app.HidePort = true

	return app.Start(fmt.Sprintf(":%d", port))
}
