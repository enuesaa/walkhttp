package main

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/enuesaa/walkhttp/pkg/cli"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/enuesaa/walkhttp/pkg/usecase"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.New()

	app := &cobra.Command{
		Use:     "walkhttp",
		Short:   "A CLI tool to call http endpoint with browser or prompt.",
		Version: "0.0.9",
		RunE: func(cmd *cobra.Command, args []string) error {
			usecase.PrintBanner(repos)

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
	app.AddCommand(cli.CtlCommand(repos))
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
