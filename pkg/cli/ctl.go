package cli

import (
	"github.com/enuesaa/walkhttp/pkg/invoke"
	"github.com/enuesaa/walkhttp/pkg/repository"
	"github.com/spf13/cobra"
)

func CtlCommand(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ctl",
		Short: "Serve web console",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			repos.Log.Printf("\n")
			repos.Log.Printf("Serving web console on localhost:%d.\n", port)
			repos.Log.Printf("\n")

			return invoke.Serve(repos, port)
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}
