package servectl

import (
	"github.com/gofiber/contrib/websocket"
)

// see https://github.com/jos-/gofiber-websocket-chat-example/blob/master/main.go
func NewWsConns() WsConns {
	return WsConns {
		conns: make(map[*websocket.Conn]int),
	}
}
type WsConns struct {
	conns map[*websocket.Conn]int
}

func (w *WsConns) Add(conn *websocket.Conn) {
	w.conns[conn] = 0
}

func (w *WsConns) Remove(conn *websocket.Conn) {
	delete(w.conns, conn)
	conn.Close()
}

func (w *WsConns) List() []*websocket.Conn {
	list := make([]*websocket.Conn, 0)
	for conn := range w.conns {
		list = append(list, conn)
	}
	return list
}
