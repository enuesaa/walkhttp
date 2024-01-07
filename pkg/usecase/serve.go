package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/pages"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

func Serve(repos repository.Repos, workdir string) error {
	pagesSrv := pages.NewPagesSrv(repos)
	pagesSrv.SetWorkdir(workdir)

	pages, err := pagesSrv.ListPages()
	if err != nil {
		return err
	}

	app := fiber.New()
	for _, page := range pages {
		route := fmt.Sprintf("/%s", page.Filename)
		app.Get(route, func(c *fiber.Ctx) error {
			invokeSrv := invoke.NewInvokeSrv(repos)
			req := invoke.Request{
				Method: page.Method,
				Url:    page.Url,
			}
			res, err := invokeSrv.Invoke(req)
			if err != nil {
				return err
			}
			return c.SendString(string(res.Body))
		})
	}

	if err := app.Listen(":3000"); err != nil {
		return err
	}
	return nil
}
