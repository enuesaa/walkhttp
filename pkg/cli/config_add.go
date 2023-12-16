package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateConfigAddCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add new path to config file.",
	}

	return cmd
}
