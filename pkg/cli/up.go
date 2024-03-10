package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/conf"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/cobra"
)

func NewMessages() Messages {
	return Messages{
		items: make([]string, 0),
	}
}

type Messages struct {
	items []string
}
func (m *Messages) Add(item string) {
	m.items = append(m.items, item)
}
func (m *Messages) First() ([]byte, error) {
	if len(m.items) > 0 {
		item := m.items[0]
		m.items = m.items[1:]
		return []byte(item), nil
	}
	return make([]byte, 0), fmt.Errorf("no item")
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

			app := fiber.New()
			app.Use(logger.New())

			messages := NewMessages()
			app.Get("/ws", websocket.New(func(c *websocket.Conn) {
				defer func ()  {
					c.Close()
					log.Println("closed")					
				}()

				for {
					time.Sleep(2 * time.Second)
					item, err := messages.First()
					if err != nil {
						continue
					}
					c.WriteMessage(websocket.TextMessage, item)
				}
			}))
			app.All("/api/*", func(c *fiber.Ctx) error {
				path := fmt.Sprintf("%s%s", config.BaseUrl, c.OriginalURL())
				messages.Add(path)
				return proxy.Do(c, path)
			})
			app.All("/*", ctlweb.Serve)
	
			return app.Listen(":3000")
		},
	}

	return cmd
}