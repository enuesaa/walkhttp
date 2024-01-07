package cli

import (
	"log"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
)

func CreateValidateCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate config files.",
		Run: func(cmd *cobra.Command, args []string) {
			taskfile, err := usecase.ParseTaskfile(repos)
			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			pp.Print(taskfile)
		},
	}

	return cmd
}

