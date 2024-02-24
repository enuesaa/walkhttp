package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)
func CreateInitCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		RunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CreateConfig(repos)
		},
	}

	return cmd
}