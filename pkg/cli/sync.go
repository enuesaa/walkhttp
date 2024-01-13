package cli

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateSyncCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "experimental sync command",
		Run: func(cmd *cobra.Command, args []string) {
			isServe, _ := cmd.Flags().GetBool("serve")
			isClient, _ := cmd.Flags().GetBool("client")
			url, _ := cmd.Flags().GetString("url")

			type Aaa struct {
				Name string `json:"name"`
				Content string `json:"content"`
			}

			if isServe {
				app := fiber.New()
				app.Get("/aaa", func(c *fiber.Ctx) error {
					res := Aaa {
						Name: "hey",
						Content: "hello",
					}
					return c.JSON(res)
				})
				if err := app.Listen("localhost:3001"); err != nil {
					log.Fatalf("Error: %s", err.Error())
				}
				return
			}

			if isClient {
				res, err := http.Get(url)
				if err != nil {
					log.Fatalf("Error: %s", err.Error())
				}
				defer res.Body.Close()
				fmt.Printf("status: %s", res.Status)
				bodybyte, err := io.ReadAll(res.Body)
				repos := repository.NewRepos()
				if err := repos.Fs.Create("aaa.json", bodybyte); err != nil {
					log.Fatalf("Error: %s", err.Error())
				}
				return
			}
		},
	}
	cmd.Flags().Bool("serve", false, "serve")
	cmd.Flags().Bool("client", false, "client")
	cmd.Flags().String("url", "", "url")

	return cmd
}
