package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePathsCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "paths",
		Short: "paths",
	}
	cmd.AddCommand(CreatePathsAddCmd(repos))
	cmd.AddCommand(CreatePathsAddOriginRequestHeaderCmd(repos))

	return cmd
}
