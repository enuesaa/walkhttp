package cli

import (
	"github.com/enuesaa/walkin/internal/handler"
	"github.com/spf13/cobra"
)

func CreateApp() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "walkin",
		Short:   "Instant web server.",
		Version: "0.0.1",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			handler.Handle(port)
		},
	}
	cmd.Flags().String("port", "3000", "port")

	// disable default
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.CompletionOptions.DisableDefaultCmd = true

	return cmd
}
