package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateValidateCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate config files.",
		Run: func(cmd *cobra.Command, args []string) {
			taskfile, err := usecase.ParseTaskfile(repos)
			if err != nil {
				repos.Log.Fatalf("Error: %s", err.Error())
			}
			repos.Log.Printf("parsed: %+v", taskfile)
		},
	}

	return cmd
}

