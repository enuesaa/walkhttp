package web

import (
	"fmt"
	"io"
	"os"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/enuesaa/walkin/internal/repository"
)

type WebService struct {
	repos repository.Repos
	port int
	serveConfig ServeConfig
}

func NewWebService(repos repository.Repos) WebService {
	return WebService{
		repos: repos,
		port: 3000,
		serveConfig: *new(ServeConfig),
	}
}

func (srv *WebService) WithPort(port int) {
	srv.port = port
}

func (srv *WebService) WithServeConfig(config ServeConfig) {
	srv.serveConfig = config
}

func (srv *WebService) calcAddress() string {
	return fmt.Sprintf(":%d", srv.port)
}

func (srv *WebService) Serve() {
	app := fiber.New()

	for path, behavior := range srv.serveConfig.Paths {
		if behavior.Behavior == Proxy {
			app.Get(path, proxy.Forward(behavior.ProxyConfig.Url))
		}
		if behavior.Behavior == ReadLocalFiles {
			app.Get(path, func(c *fiber.Ctx) error {
				requestPath := c.Path() // like `/`
		
				path := fmt.Sprintf(".%s", requestPath)
				if path == "./" {
					path = "./index.html"
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
		}
	}

	app.Listen(srv.calcAddress())
}
