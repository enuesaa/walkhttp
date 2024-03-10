package servectl

import (
	"fmt"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var broadcast = make(chan MessageBox)

//TODO: refactor
type Servectl struct {
	Wsconns WsConns
}

func (s *Servectl) Listen(port int) error {
	app := fiber.New()
	app.Use(logger.New())

	go func ()  {
		for {
			messagebox, ok := <-broadcast
			if !ok {
				return
			}
			messagebox.Send()
		}	
	}()

	app.Get("/ws", s.createHandleWs())
	app.All("/api/*", s.createHandleApi())
	app.All("/*", ctlweb.Serve)

	addr := fmt.Sprintf(":%d", port)
	return app.Listen(addr)
}
