package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/enuesaa/walkin/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/cobra"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return ctx.Status(code).JSON(ErrorResponse { Error: "Internal Sever Error" })
}

func main() {
	cmd := &cobra.Command{
		Use:     "walkin",
		Short:   "Instant web server.",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			workdir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("running on %s\n", workdir)
				
			app := fiber.New(fiber.Config{
				ErrorHandler: errorHandler,
			})
			app.Use(recover.New())
		
			queryCtl := controller.QueryController{}
			app.Post("/query", queryCtl.Query)
		
			app.Get("/", func(c *fiber.Ctx) error {
				f, err := os.Open("./testdata/index.html")
				if err != nil {
					return err
				}
				defer f.Close()
		
				content, err := io.ReadAll(f)
				if err != nil {
					return err
				}
				c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		
				return c.SendString(string(content))
			})
		
			app.Listen(":3000")
		},
	}

	// disable default
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.Execute()
}
