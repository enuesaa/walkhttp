package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateHistoryCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history",
		Short: "lookup histories",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}