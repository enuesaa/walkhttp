package cli

import (
	"log"

	"github.com/enuesaa/walkin/pkg/config"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

type ConfigFileResource struct {
	Url string `json:"url"`
}

type ConfigFile struct {
	Resources map[string]ConfigFileResource
}

func CreateConfigInitCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "create config file with prompt",
		Run: func(cmd *cobra.Command, args []string) {
			if err := config.WriteConfig(repos, config.ConfigFile{}); err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
		},
	}

	return cmd
}