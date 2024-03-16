package serve

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (s *Servectl) createHandleWs() func(*fiber.Ctx) error {
	handler := websocket.New(func(c *websocket.Conn) {
		connId := s.wsconns.Add(c)
		defer s.wsconns.Remove(connId)

		for {
			messageType, _, err := c.ReadMessage()
			if err != nil {
				break
			}
			if messageType == websocket.CloseGoingAway {
				break
			}
		}
	})

	return handler
}
