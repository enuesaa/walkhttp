package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/conf"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/cobra"
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
	for conn, _ := range w.conns {
		list = append(list, conn)
	}
	return list
}

func NewMessageBox(message string, wsconns WsConns) MessageBox {
	return MessageBox{
		message: message,
		wsconns: wsconns,
	}
}

type MessageBox struct {
	message string
	wsconns WsConns
}
func (m *MessageBox) Send() {
	for _, conn := range m.wsconns.List() {
		conn.WriteMessage(websocket.TextMessage, []byte(m.message))
	}
}

func CreateUpCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "up",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := conf.Read(repos)
			if err != nil {
				return err
			}

			wsconns := NewWsConns()
			broadcast := make(chan MessageBox)

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

			app.Get("/ws", websocket.New(func(c *websocket.Conn) {
				defer wsconns.Remove(c)
				wsconns.Add(c)

				for {
					messageType, _, err := c.ReadMessage()
					if err != nil {
						break
					}
					if messageType == websocket.CloseGoingAway {
						break
					}
				}
			}))
			app.All("/api/*", func(c *fiber.Ctx) error {
				path := fmt.Sprintf("%s%s", config.BaseUrl, c.OriginalURL())
				broadcast <- NewMessageBox(path, wsconns)
				return proxy.Do(c, path)
			})
			app.All("/*", ctlweb.Serve)
	
			return app.Listen(":3000")
		},
	}

	return cmd
}