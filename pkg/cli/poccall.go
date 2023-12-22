package cli

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreatePocCallCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poccall",
		Short: "poccall",
		Run: func(cmd *cobra.Command, args []string) {

			var message string
			type TriggerResult struct {
				Message string `json:"message"`
			}
			app := fiber.New()
			app.Post("/api/trigger", func(c *fiber.Ctx) error {
				// invoke command on handler

				command := exec.Command("curl", "-X", "POST",  "http://localhost:3000/api/callback")
				// does not wait 
				if err := command.Start(); err != nil {
					return err
				}
				for i := 0; i < 30; i++ {
					fmt.Printf("waiting %d seconds", i)
					if message == "" {
						time.Sleep(time.Second * 1)
					} else {
						break
					}
				}
				return c.JSON(TriggerResult{ Message: message })
			})
			app.Post("/api/callback", func(c *fiber.Ctx) error {
				fmt.Printf("callback called\n")
				message = "hello"
				return nil
			})
			if err := app.Listen(":3000"); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}


		},
	}

	return cmd
}