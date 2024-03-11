package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateUpCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "up",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			return usecase.Serve(port)
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}
