package serve

import (
	"fmt"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() Servectl {
	return Servectl{
		wsconns: make(map[*websocket.Conn]int),
		wssend:  make(chan []byte),
		Port:    3000,
	}
}

type Servectl struct {
	wsconns map[*websocket.Conn]int
	wssend  chan []byte
	Port    int
}

func (s *Servectl) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s *Servectl) Listen() error {
	app := fiber.New()
	app.Use(logger.New())

	go s.broadcast()
	app.Get("/ws", s.createHandleWs())
	app.All("/api/*", s.handleApi)
	app.All("/*", ctlweb.Serve)

	return app.Listen(s.Addr())
}
