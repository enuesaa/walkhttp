package serve

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

func NewWsConns() WsConns {
	return WsConns {
		conns: make(map[string]*websocket.Conn),
	}
}

type WsConns struct {
	conns map[string]*websocket.Conn
}

func (c *WsConns) Add(conn *websocket.Conn) string {
	id := uuid.NewString()
	c.conns[id] = conn
	return id
}

func (c *WsConns) Remove(id string) {
	if conn, ok := c.conns[id]; ok {
		delete(c.conns, id)
		conn.Close()
	}
}

func (c *WsConns) Send(message string) {
	for id, conn := range c.conns {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			c.Remove(id)
		}
	}
}
