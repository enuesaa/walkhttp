package routeproxy

import (
	"fmt"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/labstack/echo/v4"
)

func Handle(repos repository.Repos) echo.HandlerFunc {
	handler := func(c echo.Context) error {
		invokeSrv := invoke.New(repos)

		// todo
		url := fmt.Sprintf("https://example.com%s", c.Request().URL.Path)

		entry := invoke.Entry{
			Id: "",
			Request: invoke.EntryRequest{
				Method:  c.Request().Method,
				Url:     url,
				Headers: map[string][]string{},
				Body:    "",
				Started: 0,
			},
			Response: invoke.EntryResponse{
				Status:   0,
				Headers:  map[string][]string{},
				Body:     "",
				Received: 0,
			},
		}
	
		// for _, h := range c.Request().Header {
		// 	if _, ok := entry.Request.Headers[h.Name]; !ok {
		// 		entry.Request.Headers[h.Name] = []string{}
		// 	}
		// 	entry.Request.Headers[h.Name] = append(entry.Request.Headers[h.Name], h.Value)
		// }
	
		if err := invokeSrv.Invoke(&entry); err != nil {
			return err
		}
		if err := invokeSrv.Write(entry); err != nil {
			return err
		}
		return nil
	}

	return handler
}
