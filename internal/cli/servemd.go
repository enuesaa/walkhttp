package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark-meta"
)

type ServemdResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	Content string `json:"content"`
}

func CreateServemdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve-md",
		Short: "serve markdown",
		Run: func(cmd *cobra.Command, args []string) {
			app := fiber.New()
			app.Get("/", func(c *fiber.Ctx) error {

				f, err := os.Open("testdata/example.md")
				if err != nil {
					return err
				}
				defer f.Close()
				source, err := io.ReadAll(f)
				if err != nil {
					return err
				}
				var parsed bytes.Buffer
				markdown := goldmark.New(goldmark.WithExtensions(meta.Meta))
				context := parser.NewContext()
				if err := markdown.Convert(source, &parsed, parser.WithContext(context)); err != nil {
					return err
				}
				// see https://github.com/yuin/goldmark-meta
				metadata := meta.Get(context)
				fmt.Println(metadata)

				res := ServemdResponse{
					Title: metadata["title"].(string),
					Description: metadata["description"].(string),
					// Tags: metadata["tags"].([]string),
					Content: parsed.String(),
				}

				return c.JSON(res)
			})
			app.Listen(":3000")
		},
	}

	return cmd
}