package cli

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreatePocCallCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poccall",
		Short: "poccall",
		Run: func(cmd *cobra.Command, args []string) {

			type TriggerResult struct {
				Message string `json:"message"`
			}
			app := fiber.New()
			app.Post("/api/trigger", func(c *fiber.Ctx) error {
				// invoke command on handler

				command := exec.Command("printf", `{"message": "%s"}`, "hello")
				result, err := command.Output()
				if err != nil {
					return err
				}
				var message TriggerResult
				json.Unmarshal(result, &message)
				return c.JSON(message)
			})
			if err := app.Listen(":3000"); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}


		},
	}

	return cmd
}