package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateCtlCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ctl",
		Short: "ctl",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			if err := usecase.CreateWalkindir(repos); err != nil {
				return err
			}

			return usecase.Serve(repos, port)
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}
