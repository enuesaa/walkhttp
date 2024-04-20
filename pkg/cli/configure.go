package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateConfigureCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "configure",
		RunE: func(cmd *cobra.Command, args []string) error {
			// configure global settings.
			return nil
		},
	}

	return cmd
}
