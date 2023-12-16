package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateConfigCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "configure",
	}
	cmd.AddCommand(CreateConfigInitCmd(repos))
	cmd.AddCommand(CreateConfigAddCmd(repos))

	return cmd
}