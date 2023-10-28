package main

import (
	"github.com/enuesaa/walkin/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use:     "walkin",
		Short:   "Instant web server.",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	app.AddCommand(cli.CreateServeCmd())
	app.AddCommand(cli.CreateServemdCmd())

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.CompletionOptions.DisableDefaultCmd = true

	app.Execute()
}
