package cli

import (
	"github.com/spf13/cobra"
)

func CreateApp() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "walkin",
		Short:   "Instant web server.",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(CreateServeCmd())
	cmd.AddCommand(CreateServemdCmd())
	cmd.AddCommand(CreateCreateConfigCmd())

	// disable default
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.CompletionOptions.DisableDefaultCmd = true

	return cmd
}
