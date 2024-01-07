package cli

import (
	"log"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateRunCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run",
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.Invoke(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}

	return cmd
}

