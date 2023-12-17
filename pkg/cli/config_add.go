package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/walkin/pkg/config"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateConfigAddCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add new path to config file.",
		Run: func(cmd *cobra.Command, args []string) {
			configfile, err := config.ReadConfig(repos)
			if err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
			fmt.Printf("%+v\n", configfile)
		},
	}

	return cmd
}
