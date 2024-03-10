package serve

import (
	"github.com/gofiber/contrib/websocket"
)

// see https://github.com/jos-/gofiber-websocket-chat-example/blob/master/main.go

func (s *Servectl) addConn(conn *websocket.Conn) {
	s.wsconns[conn] = 0
}
func (s *Servectl) rmConn(conn *websocket.Conn) {
	delete(s.wsconns, conn)
	conn.Close()
}

func (s *Servectl) listConn() []*websocket.Conn {
	list := make([]*websocket.Conn, 0)
	for conn := range s.wsconns {
		list = append(list, conn)
	}
	return list
}
