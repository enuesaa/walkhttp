package cli

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve static files",
		Run: func(cmd *cobra.Command, args []string) {		
			app := fiber.New()
			app.Get("/*", func(c *fiber.Ctx) error {
				requestPath := c.Path() // like `/`

				path := fmt.Sprintf(".%s", requestPath) // like `./`
				if strings.HasSuffix(path, "/") {
					path += "index.html"
				}

				f, err := os.Open(path)
				if err != nil {
					return err
				}
				defer f.Close()

				content, err := io.ReadAll(f)
				if err != nil {
					return err
				}
				mimeType := http.DetectContentType(content)
				c.Set(fiber.HeaderContentType, mimeType)

				return c.SendString(string(content))
			})
			app.Listen(":3000")
		},
	}

	return cmd
}
