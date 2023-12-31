package cli

import (
	"log"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

type TriggerResult struct {
	Message string `json:"message"`
}

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			workdir, _ := cmd.Flags().GetString("workdir")

			if err := usecase.Serve(repos, workdir); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
		},
	}
	cmd.Flags().String("workdir", ".", "workdir")

	return cmd
}

