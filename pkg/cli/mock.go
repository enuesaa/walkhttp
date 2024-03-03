package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateMockCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mock",
		Short: "mock",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}