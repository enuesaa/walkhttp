package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
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

			app := fiber.New()
			app.Post("/api/trigger", func(c *fiber.Ctx) error {
				command := exec.Command("printf", `{"message": "%s"}`, "hello")
				result, err := command.Output()
				if err != nil {
					return err
				}
				var message TriggerResult
				json.Unmarshal(result, &message)

				invokeSrv := invoke.NewInvokeSrv(repos)
				res, err := invokeSrv.Invoke("GET", "http://example.com")
				if err != nil {
					log.Fatalf("Error: %s\n", err)
				}
				fmt.Printf("%s\n", string(res))
				return nil
			})
			if err := app.Listen(":3000"); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
		},
	}

	return cmd
}
