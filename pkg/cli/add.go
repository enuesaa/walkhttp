package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/walkin/pkg/endpoint"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateAddCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			fmt.Printf("%s\n", name)

			path, _ := cmd.Flags().GetString("path")
			url, _ := cmd.Flags().GetString("url")

			configSrv := endpoint.NewConfigSrv(repos)
			configjson, err := configSrv.Read()
			if err != nil {
				configjson = endpoint.Config{
					Paths: make([]endpoint.ConfigPath, 0),
				}
			}
			configjson.Paths = append(configjson.Paths, endpoint.ConfigPath{
				Path: path,
				Url: url,
			})

			if err := configSrv.Write(configjson); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().String("name", "", "name")
	cmd.Flags().String("path", "/", "proxy path")
	cmd.Flags().String("url", "https://example.com/", "proxy url")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("path")
	cmd.MarkFlagRequired("url")

	return cmd
}
