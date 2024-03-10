package cli

import (
	"fmt"
	"log"

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

var broadcast = make(chan string)

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

			app.Get("/ws", websocket.New(func(c *websocket.Conn) {
				defer func() {
					log.Printf("close\n")
					c.Close()
				}()

				go func ()  {
					for {
						message, ok := <-broadcast
						if !ok {
							return
						}
						log.Printf("a: %s\n", message)
						if err := c.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
							log.Printf("err: %s\n", err.Error())
						}
					}	
				}()
				for {
					messageType, _, err := c.ReadMessage()
					if err != nil {
						log.Printf("err: %s\n", err.Error())
						break
					}
					if messageType == websocket.CloseGoingAway {
						log.Printf("closed\n")
					}
				}
			}))
			app.All("/api/*", func(c *fiber.Ctx) error {
				path := fmt.Sprintf("%s%s", config.BaseUrl, c.OriginalURL())
				log.Println(path)
				broadcast <- path					
				log.Println(path)
				return proxy.Do(c, path)
			})
			app.All("/*", ctlweb.Serve)
	
			return app.Listen(":3000")
		},
	}

	return cmd
}