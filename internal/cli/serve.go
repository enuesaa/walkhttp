package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve instant web server.",
		Run: func(cmd *cobra.Command, args []string) {
			workdir, _ := os.Getwd()
			fmt.Printf("running on %s\n", workdir)
		
			app := fiber.New()
			app.Get("/*", func(c *fiber.Ctx) error {
				path := c.OriginalURL() // like `/`

				f, err := os.Open(fmt.Sprintf(".%s", path))
				if err != nil {
					return err
				}
				defer f.Close()

				content, err := io.ReadAll(f)
				if err != nil {
					return err
				}
				// c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

				return c.SendString(string(content))
			})
			app.Listen(":3000")
		},
	}
	cmd.Flags().String("config", "config.json", "config file path")

	return cmd
}
