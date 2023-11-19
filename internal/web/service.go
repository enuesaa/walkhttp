package web

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/enuesaa/walkin/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
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

	fmt.Printf("%+v\n", srv.serveConfig)

	for path, behavior := range srv.serveConfig.Paths {
		if behavior.Behavior == Proxy {
			// TODO check
			if path == "" {
				continue
			}
			app.Get(path, createProxyHandler(path, behavior.ProxyConfig))
		}
		if behavior.Behavior == ReadLocalFiles {
			app.Get(path, createReadLocalFilesHandler(path))	
		}
	}

	app.Listen(srv.calcAddress())
}

func createProxyHandler(path string, proxyConfig ProxyConfig) func(c *fiber.Ctx) error {
	fmt.Printf("[walkin] %v\n", proxyConfig)
	fmt.Printf("[walkin] proxy on %s%s\n", path, proxyConfig.Url)
	return proxy.Forward(proxyConfig.Url)
}

func createReadLocalFilesHandler(path string) func(c *fiber.Ctx) error {
	fmt.Printf("[walkin] readLocalFiles on %s\n", path)

	handler := func(c *fiber.Ctx) error {
		requestPath := c.Path() // like `/`

		basePath := strings.Replace(path, "*", "", -1)
		lookupPath := fmt.Sprintf("./%s", strings.TrimPrefix(requestPath, basePath))

		if strings.HasSuffix(lookupPath, "/") {
			lookupPath = lookupPath + "index.html"
		}
		fmt.Printf("lookupPath is %s\n", lookupPath)

		f, err := os.Open(lookupPath)
		if err != nil {
			return err
		}
		defer f.Close()

		content, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		return c.SendString(string(content))
	}

	return handler
}
