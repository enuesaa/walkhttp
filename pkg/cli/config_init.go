package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateConfigInitCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "create config file with prompt",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}