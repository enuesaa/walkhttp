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

			value, err := repos.Prompt.Ask("Aaa: ", "")
			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			fmt.Printf("%s\n", value)
			choice, err := repos.Prompt.Select("Aaa: ", []string{"header", "body", "query"})
			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			fmt.Printf("%s\n", choice)
		},
	}

	return cmd
}