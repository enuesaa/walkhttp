package serve

import (
	"github.com/gofiber/contrib/websocket"
)

func (s *Servectl) broadcast() {
	for {
		data, ok := <-s.wssend
		if !ok {
			return
		}
		for _, conn := range s.listConn() {
			if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
				s.rmConn(conn)
			}
		}
	}
}
