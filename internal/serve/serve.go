package serve

import (
	// "net/url"

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
	app.Any("/graphql", ctl.handleGql())
	app.GET("/graphql/playground", ctl.handleGqlPlayground())
	app.Any("/*", ctlweb.Serve)
	app.Any("/", ctlweb.Serve)

	// proxy
	// proxyroot := app.Group("/proxy")
	// exampleSite, err := url.Parse("https://example.com")
	// if err != nil {
	// 	return err
	// }
	// proxyroot.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
	// 	Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
	// 		{URL: exampleSite},
	// 	}),
	// }))

	app.HideBanner = true
	app.HidePort = true

	return app.Start(ctl.Address())
}
