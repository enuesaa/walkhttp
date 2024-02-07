package cli

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve",
		Run: func(cmd *cobra.Command, args []string) {	
			app := fiber.New()

			app.Get("/*", func(c *fiber.Ctx) error {
				path := c.Path() // like `/`
				if strings.HasSuffix(path, "/") {
					path = filepath.Join(path, "index.html")					
				}
				if !strings.Contains(path, ".") {
					path = path + ".html"
				}
				// TODO remove
				path = strings.ReplaceAll(path, "%5B", "[")
				path = strings.ReplaceAll(path, "%5D", "]")
				path = "." + path 
				fmt.Printf("%s\n", path)

				f, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				fileExt := filepath.Ext(path)
				mimeType := mime.TypeByExtension(fileExt)
				c.Set(fiber.HeaderContentType, mimeType)

				return c.SendString(string(f))
			})
			app.Listen(":3000")
		},
	}

	return cmd
}
