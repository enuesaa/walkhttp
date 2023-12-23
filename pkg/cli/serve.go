package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

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

				client := &http.Client{}

				req, err := http.NewRequest("GET", "http://example.com", nil)
				if err != nil {
					return err
				}
				req.Header.Add("X-Message", message.Message)
			
				res, err := client.Do(req)
				if err != nil {
					log.Fatalf("Error: %s\n", err)
				}
				defer res.Body.Close()
				resbodybytes, err := io.ReadAll(res.Body)
				if err != nil {
					log.Fatalf("Error: %s\n", err)
				}
				fmt.Printf("%s\n", string(resbodybytes))
				return nil
			})
			if err := app.Listen(":3000"); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
		},
	}

	return cmd
}
