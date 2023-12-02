package main

import (
	"github.com/enuesaa/walkin/internal/cli"
	"github.com/enuesaa/walkin/internal/repository"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.NewRepos()

	app := &cobra.Command{
		Use:     "walkin",
		Short:   "Instant web server.",
		Version: "0.0.1",
	}
	app.AddCommand(cli.CreateServeCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.CompletionOptions.DisableDefaultCmd = true

	app.Execute()
}
