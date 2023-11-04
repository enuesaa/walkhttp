package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type WebService struct {
	api bool
	staticServe bool
	graphql bool
	port int
}

func NewWebService() WebService {
	return WebService{
		api: false,
		staticServe: false,
		graphql: false,
		port: 3000,
	}
}

type ApiRespponseFile struct {
	Name string `json:"name"`
	IsDir bool `json:"isDir"`
}

type ApiResponse struct {
	Files []ApiRespponseFile `json:"files"`
}

func (srv *WebService) EnableApi() {
	srv.api = true
}

func (srv *WebService) EnableStaticServe() {
	srv.staticServe = true
}

func (srv *WebService) EnableGraphql() {
	srv.graphql = true
}

func (srv *WebService) SetPort(port int) {
	srv.port = port
}

func (srv *WebService) Serve() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	if srv.staticServe {
		app.Get("/*", func(c *fiber.Ctx) error {
			requestPath := c.Path() // like `/`
	
			path := fmt.Sprintf(".%s", requestPath) // like `./`
			if strings.HasSuffix(path, "/") {
				path += "index.html"
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
			mimeType := http.DetectContentType(content)
			c.Set(fiber.HeaderContentType, mimeType)
			fmt.Printf("%s is %s.", path, mimeType)
	
			return c.SendString(string(content))
		})
	}
	
	if srv.api {
		app.Get("/", func(c *fiber.Ctx) error {
			files, err := os.ReadDir("./")
			if err != nil {
				return err
			}

			res := ApiResponse {
				Files: make([]ApiRespponseFile, 0),
			}
			for _, f := range files {
				res.Files = append(res.Files, ApiRespponseFile{
					Name: f.Name(),
					IsDir: f.IsDir(),
				})
			}

			return c.JSON(res)
		})

	}

	addr := fmt.Sprintf(":%d", srv.port)
	fmt.Printf("serve on http://localhost%s\n", addr)
	app.Listen(addr)
}
