package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/conf"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/cobra"
)

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
		
			app.Get("/*", func(c *fiber.Ctx) error {
				path := config.BaseUrl + c.OriginalURL()
				fmt.Printf("path: %s\n", path)
				return proxy.Do(c, path)
			})
		
			return app.Listen(":3000")
		},
	}

	return cmd
}