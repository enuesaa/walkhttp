package serve

import (
	"fmt"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(repos repository.Repos) Servectl {
	return Servectl{
		repos:   repos,
		wsconns: NewWsConns(),
		Port:    3000,
	}
}

type Servectl struct {
	repos   repository.Repos
	wsconns WsConns
	Port    int
}

func (s *Servectl) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *Servectl) Listen() error {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/ws", s.createHandleWs())
	app.All("/api/*", s.handleApi)
	app.All("/*", ctlweb.Serve)

	return app.Listen(s.Addr())
}
