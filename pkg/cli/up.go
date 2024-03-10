package cli

import (
	"log"
	"time"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

type Messages struct {
	Items []string
}

func CreateUpCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "up",
		RunE: func(cmd *cobra.Command, args []string) error {
			// config, err := conf.Read(repos)
			// if err != nil {
			// 	return err
			// }

			app := fiber.New()
			app.Use(logger.New())

			messages := Messages{
				Items: make([]string, 0),
			}

			app.Get("/ws", websocket.New(func(c *websocket.Conn) {
				defer func ()  {
					c.Close()
					log.Println("closed")					
				}()

				for range 10 {
					time.Sleep(2 * time.Second)
					if len(messages.Items) > 0 {
						c.WriteMessage(websocket.TextMessage, []byte(messages.Items[0]))
						messages.Items = messages.Items[1:]
					}
				}
			}))

			app.All("/api/*", func(c *fiber.Ctx) error {
				messages.Items = append(messages.Items, c.Path())
				return nil
			})

			app.All("/*", ctlweb.Serve)
	
			// app.Get("/*", func(c *fiber.Ctx) error {
			// 	path := config.BaseUrl + c.OriginalURL()
			// 	fmt.Printf("path: %s\n", path)
			// 	return proxy.Do(c, path)
			// })
		
			return app.Listen(":3000")
		},
	}

	return cmd
}