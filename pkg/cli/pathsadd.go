package cli

import (
	"log"

	"github.com/enuesaa/walkin/pkg/config"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePathsAddCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			url, _ := cmd.Flags().GetString("url")

			configSrv := config.NewConfigSrv(repos)
			configjson, err := configSrv.Read()
			if err != nil {
				configjson = config.Config{
					Paths: make([]config.ConfigPath, 0),
				}
			}
			configjson.Paths = append(configjson.Paths, config.ConfigPath{
				Path: path,
				Url: url,
			})

			if err := configSrv.Write(configjson); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().String("path", "/", "proxy path")
	cmd.Flags().String("url", "https://example.com/", "proxy url")

	cmd.MarkFlagRequired("path")
	cmd.MarkFlagRequired("url")

	return cmd
}

