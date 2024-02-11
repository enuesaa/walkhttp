package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateOptionsCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "options",
		Short: "options request",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}