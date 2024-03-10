package servectl

import (
	"github.com/gofiber/contrib/websocket"
)

func NewMessageBox(message []byte, wsconns *WsConns) MessageBox {
	return MessageBox{
		message: message,
		wsconns: wsconns,
	}
}

type MessageBox struct {
	message []byte
	wsconns *WsConns
}

func (m *MessageBox) Send() {
	for _, conn := range m.wsconns.List() {
		conn.WriteMessage(websocket.TextMessage, m.message)
	}
}
