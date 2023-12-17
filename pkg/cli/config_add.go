package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateConfigAddCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add new path to config file.",
		Run: func(cmd *cobra.Command, args []string) {
			fbyte, err := repos.Fs.Read("config.json")
			if err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}

			var config ConfigFile
			if err := json.Unmarshal(fbyte, &config); err != nil {
				log.Fatalf("Errpr: %s\n", err.Error())
			}

			fmt.Printf("%+v\n", config)
		},
	}

	return cmd
}
