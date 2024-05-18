package cli

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateCtlCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ctl",
		Short: "ctl",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			return invoke.ServeGraphql(repos, port)
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}
