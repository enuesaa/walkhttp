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
func NewWsConn() WsConn {
	return WsConn {
		conns: make(map[*websocket.Conn]int),
		message: "",
	}
}
type WsConn struct {
	conns map[*websocket.Conn]int
	message string
}
func (w *WsConn) Add(conn *websocket.Conn) {
	w.conns[conn] = 0
}
func (w *WsConn) Remove(conn *websocket.Conn) {
	delete(w.conns, conn)
	conn.Close()
}
func (w *WsConn) WithMessage(message string) {
	w.message = message
}
func (w *WsConn) Send() {
	for conn, _ := range w.conns {
		conn.WriteMessage(websocket.TextMessage, []byte(w.message))
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

			wsconn := NewWsConn()
			broadcast := make(chan WsConn)

			app := fiber.New()
			app.Use(logger.New())

			go func ()  {
				for {
					wsconn, ok := <-broadcast
					if !ok {
						return
					}
					wsconn.Send()
				}	
			}()

			app.Get("/ws", websocket.New(func(c *websocket.Conn) {
				defer wsconn.Remove(c)
				wsconn.Add(c)

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
				wsconn.WithMessage(path)
				broadcast <- wsconn
				return proxy.Do(c, path)
			})
			app.All("/*", ctlweb.Serve)
	
			return app.Listen(":3000")
		},
	}

	return cmd
}