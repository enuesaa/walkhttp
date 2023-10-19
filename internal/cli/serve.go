package cli

import (
	_ "github.com/enuesaa/walkin/internal/handler"
	"github.com/spf13/cobra"
)

func CreateServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve instant web server.",
		Run: func(cmd *cobra.Command, args []string) {
			// config, _ := cmd.Flags().GetString("config")
			// handler.Handle(port)
		},
	}
	cmd.Flags().String("config", "config.json", "config file path")

	return cmd
}
