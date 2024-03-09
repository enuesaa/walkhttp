package cli

import (
	"log"

	"github.com/enuesaa/walkin/ctlweb"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

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

			// see: https://note.com/twsnmp/n/ne64357e08038
			// 記事にある通り hub の方が良さそう
			app.Get("/ws", websocket.New(func(c *websocket.Conn) {
				// これ要は開きっぱなしで占有している
				for {
					mtype, msg, err := c.ReadMessage()
					if err != nil {
						break
					}
					log.Printf("Read: %s", msg)
					err = c.WriteMessage(mtype, msg)
					if err != nil {
						log.Println(err)
						break
					}
				}
				log.Println("closed")
			}))

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