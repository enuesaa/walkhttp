package cli

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/enuesaa/walkin/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateMockCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mock",
		Short: "mock",
		Run: func(cmd *cobra.Command, args []string) {	
			app := fiber.New()
			app.Get("/*", func(c *fiber.Ctx) error {
				path := c.Path() // like `/`
				path = fmt.Sprintf("dist%s", path)
				if strings.HasSuffix(path, "/") {
					path += "index.html"
				}

				// read file here.

				fileExt := filepath.Ext(path)
				mimeType := mime.TypeByExtension(fileExt)
				c.Set(fiber.HeaderContentType, mimeType)

				return c.SendString("")
			})

			app.Listen(":3000")
		},
	}

	return cmd
}
