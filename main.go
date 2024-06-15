package main

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.New()

	app := &cobra.Command{
		Use:     "walkin",
		Short:   "A CLI tool to call http endpoint with browser or prompt.",
		Version: "0.0.7",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")
			repos.Log.Printf("\n")
			repos.Log.Printf("Serving web console on localhost:%d.\n", port)
			repos.Log.Printf("You can also call http endpoint via prompt.\n")
			repos.Log.Printf("\n")

			go invoke.Serve(repos, port)

			for {
				if err := usecase.Prompt(repos); err != nil {
					repos.Log.Printf("Error: %s", err.Error())
					if err == huh.ErrUserAborted {
						break
					}
				}
			}

			return nil
		},
	}
	app.Flags().Int("port", 3000, "port")

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
