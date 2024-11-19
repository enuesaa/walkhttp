package router

import (
	"github.com/enuesaa/walkhttp/ctlweb"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/routegql"
	"github.com/enuesaa/walkhttp/internal/routegqlplayground"
	"github.com/enuesaa/walkhttp/internal/routeproxy"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(repos repository.Repos) *echo.Echo {
	app := echo.New()

	// middleware
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// routes
	app.Any("/_/graphql", routegql.Handle(repos))
	app.GET("/_/graphql/playground", routegqlplayground.Handle())
	app.Any("/_", ctlweb.Handle())
	app.Any("/_/*", ctlweb.Handle())
	app.Any("/*", routeproxy.Handle(repos))

	app.HideBanner = true
	app.HidePort = true

	return app
}
