package routeproxy

import (
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

		if err := invokeSrv.Invoke(&entry); err != nil {
			return err
		}
		if err := invokeSrv.Save(entry); err != nil {
			return err
		}
		return nil
	}

	return handler
}
