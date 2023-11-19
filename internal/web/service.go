package web

import (
	"fmt"
	"io"
	"os"
	"github.com/gofiber/fiber/v2"
	"os/signal"
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

	// see https://github.com/gofiber/fiber/issues/899
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
			_ = <-c
			fmt.Println("Gracefully shutting down...")
			_ = app.Shutdown()
	}()

	app.Get("/*", func(c *fiber.Ctx) error {
		requestPath := c.Path() // like `/`

		path := fmt.Sprintf(".%s", requestPath)
		if path == "./" {
			path = "./testdata/index.html" // TODO for dev.
		}
		f, err := os.Open(path)
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
	fmt.Println("Running cleanup tasks...")
}
