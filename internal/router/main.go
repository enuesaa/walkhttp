package router

import (
	"github.com/enuesaa/walkhttp/ctlweb"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/routegql"
	"github.com/enuesaa/walkhttp/internal/routegqlplayground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(repos repository.Repos) *echo.Echo {
	app := echo.New()

	// middleware
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// routes
	app.Any("/graphql", routegql.Handle(repos))
	app.GET("/graphql/playground", routegqlplayground.Handle())
	app.Any("/ctlweb", ctlweb.Handle())
	app.Any("/ctlweb/*", ctlweb.Handle())
	// app.Any("/", ctlweb.Serve)

	app.HideBanner = true
	app.HidePort = true

	return app
}
