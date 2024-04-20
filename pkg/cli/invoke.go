package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateInvokeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke",
		RunE: func(cmd *cobra.Command, args []string) error {
			// invoke http request with template name.
			return nil
		},
	}

	return cmd
}
