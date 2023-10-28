package cli

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type ApiRespponseFile struct {
	Name string `json:"name"`
	IsDir bool `json:"isDir"`
}

type ApiResponse struct {
	Files []ApiRespponseFile `json:"files"`
}

func CreateServeApiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve-api",
		Short: "serve api",
		Run: func(cmd *cobra.Command, args []string) {	
			port, _ := cmd.Flags().GetInt("port")
	
			app := fiber.New(fiber.Config{
				DisableStartupMessage: true,
			})
			app.Get("/", func(c *fiber.Ctx) error {
				files, err := os.ReadDir("./")
				if err != nil {
					return err
				}

				res := ApiResponse {
					Files: make([]ApiRespponseFile, 0),
				}
				for _, f := range files {
					res.Files = append(res.Files, ApiRespponseFile{
						Name: f.Name(),
						IsDir: f.IsDir(),
					})
				}

				return c.JSON(res)
			})

			addr := fmt.Sprintf(":%d", port)
			fmt.Printf("serve on http://localhost%s\n", addr)
			app.Listen(addr)
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}