package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type TriggerResult struct {
	Message string `json:"message"`
}

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			workdir, _ := cmd.Flags().GetString("workdir")

			pages, err := usecase.ListPages(repos, workdir)
			if err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}

			app := fiber.New()

			for _, page := range pages {
				fmt.Printf("found: %s\n", page.Filename)

				route := fmt.Sprintf("/%s", page.Filename)
				app.Get(route, func(c *fiber.Ctx) error {
					invokeSrv := invoke.NewInvokeSrv(repos)
					req := invoke.Request{
						Method: page.Method,
						Url: page.Url,
					}
					res, err := invokeSrv.Invoke(req)
					if err != nil {
						return err
					}
					return c.SendString(string(res.Body))
				})
			}

			if err := app.Listen(":3000"); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
		},
	}
	cmd.Flags().String("workdir", ".", "workdir")

	return cmd
}

// command := exec.Command("printf", `{"message": "%s"}`, "hello")
// result, err := command.Output()
// if err != nil {
// 	return err
// }
// var message TriggerResult
// json.Unmarshal(result, &message)
