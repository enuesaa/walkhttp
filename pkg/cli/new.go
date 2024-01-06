package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateNewCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new",
		Short: "new",
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.CreateConfigFileIfNotExist(repos); err != nil {
				log.Fatalf("Error: failed to create config file")
			}

			endpointDef, err := usecase.DefineEndpointWithPrompt(repos)
			if err != nil {
				log.Fatalf("Error: failed to run prompt.")
			}
			fmt.Printf("%+v\n", endpointDef)
		},
	}

	return cmd
}