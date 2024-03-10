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

var wsconn *websocket.Conn
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

			go func ()  {
				for {
					message, ok := <-broadcast
					if !ok {
						return
					}
					log.Printf("a: %s\n", message)
					if wsconn == nil {
						return
					}
					log.Printf("b: %s\n", message)

					if err := wsconn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						wsconn.Close()
						wsconn = nil
						log.Printf("closed: %s\n", err.Error())
					}
					log.Printf("c: %s\n", message)
					return
				}	
			}()

			app.Get("/ws", websocket.New(func(c *websocket.Conn) {
				// defer func() {
				// 	log.Printf("close\n")
				// 	c.Close()
				// 	wsconn = nil
				// }()
				wsconn = c
			}))
			app.All("/api/*", func(c *fiber.Ctx) error {
				path := fmt.Sprintf("%s%s", config.BaseUrl, c.OriginalURL())
				log.Println(path)
				go func ()  {
					broadcast <- path					
				}()
				log.Println(path)
				return proxy.Do(c, path)
			})
			app.All("/*", ctlweb.Serve)
	
			return app.Listen(":3000")
		},
	}

	return cmd
}