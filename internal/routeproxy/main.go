package routeproxy

import (
	// "strings"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/labstack/echo/v4"
)

func Handle(repos repository.Repos) echo.HandlerFunc {
	handler := func(c echo.Context) error {
		invokeSrv := invoke.New(repos)

		method := c.Request().Method
		path := c.Request().URL.Path
		headers := c.Request().Header

		entry := invokeSrv.NewEntry(method, path)
		for name, values := range headers {
			entry.Request.Headers[name] = values
		}
		delete(entry.Request.Headers, "Accept-Encoding")

		if err := invokeSrv.Invoke(&entry); err != nil {
			return err
		}
		if err := invokeSrv.Save(entry); err != nil {
			return err
		}

		for name, values := range entry.Response.Headers {
			for i, value := range values {
				if i == 0 {
					c.Response().Header().Set(name, value)
					continue
				}
				c.Response().Header().Add(name, value)
			}
		}

		return c.String(entry.Response.Status, entry.Response.Body)
	}

	return handler
}
