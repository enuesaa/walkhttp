package main

import (
	"fmt"
	"log"
	"time"

	"github.com/enuesaa/walkhttp/pkg/cli"
	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.New()

	app := &cobra.Command{
		Use:     "walkhttp",
		Short:   "A CLI tool to call http endpoint with browser or prompt.",
		Version: "0.0.9",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			repos.Log.Printf("\n")
			repos.Log.Printf("Serving web console on localhost:%d.\n", port)
			repos.Log.Printf("\n")

			go func() {
				time.Sleep(1 * time.Second)
				url := fmt.Sprintf("http://localhost:%d", port)
				browser.OpenURL(url)
			}()

			return invoke.Serve(repos, port)
		},
	}
	app.Flags().Int("port", 3000, "port")

	app.AddCommand(cli.Prompt(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceErrors = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	if err := app.Execute(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
