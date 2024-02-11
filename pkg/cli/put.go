package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePutCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put",
		Short: "put request",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}