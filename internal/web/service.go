package web

import (
	"fmt"
	"io"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/enuesaa/walkin/internal/repository"
)

type WebService struct {
	repos repository.Repos
	port int
}

func NewWebService(repos repository.Repos) WebService {
	return WebService{
		repos: repos,
		port: 3000,
	}
}

func (srv *WebService) WithPort(port int) {
	srv.port = port
}

func (srv *WebService) calcAddress() string {
	return fmt.Sprintf(":%d", srv.port)
}

func (srv *WebService) Serve() {
	app := fiber.New()

	app.Get("/*", func(c *fiber.Ctx) error {
		requestPath := c.Path() // like `/`
		f, err := os.Open(fmt.Sprintf(".%s", requestPath))
		if err != nil {
			return err
		}
		defer f.Close()

		content, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		return c.SendString(string(content))
	})

	app.Listen(srv.calcAddress())
}
