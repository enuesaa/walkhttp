package serve

import (
	"github.com/enuesaa/walkhttp/ctlweb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (ctl *ServeCtl) Serve() error {
	app := echo.New()

	// middleware
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// routes
	app.Any("/graph", ctl.handleGql())
	app.GET("/graph/playground", ctl.handleGqlPlayground())
	app.Any("/*", ctlweb.Serve)
	app.Any("/", ctlweb.Serve)

	app.HideBanner = true
	app.HidePort = true

	return app.Start(ctl.Address())
}
