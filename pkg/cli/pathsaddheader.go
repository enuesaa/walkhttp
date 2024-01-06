package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/walkin/pkg/config"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePathsAddOriginRequestHeaderCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-origin-request-header",
		Short: "add header of origin request",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			name, _ := cmd.Flags().GetString("name")
			value, _ := cmd.Flags().GetString("value")

			fmt.Printf("%s: %s, %s\n", path, name, value)

			configSrv := config.NewConfigSrv(repos)
			configjson, err := configSrv.Read()
			if err != nil {
				log.Fatalf("Error: config file does not exist.")
			}

			paths := make([]config.ConfigPath, 0)
			for _, configpath := range configjson.Paths {
				if configpath.Path == path {
					headers := configpath.OriginRequestHeaders
					headers[name] = value
					paths = append(paths, config.ConfigPath{
						Path: path,
						Url: configpath.Url,
						OriginRequestHeaders: headers,
					})
				} else {
					paths = append(paths, configpath)
				}
			}

			if err := configSrv.Write(configjson); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().String("path", "/", "proxy path")
	cmd.Flags().String("name", "Accept", "header name")
	cmd.Flags().String("value", "application/json", "header value")

	cmd.MarkFlagRequired("path")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("value")

	return cmd
}