package cli

import (
	"fmt"
	// "log"

	// "github.com/enuesaa/walkin/pkg/endpoint"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateTransformRequestCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transform-request",
		Short: "transform reuqest",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			name, _ := cmd.Flags().GetString("name")
			value, _ := cmd.Flags().GetString("value")

			fmt.Printf("%s: %s, %s\n", path, name, value)

			// configSrv := endpoint.NewConfigSrv(repos)
			// configjson, err := configSrv.Read()
			// if err != nil {
			// 	log.Fatalf("Error: config file does not exist.")
			// }

			// paths := make([]endpoint.ConfigPath, 0)
			// for _, configpath := range configjson.Paths {
			// 	if configpath.Path == path {
			// 		headers := configpath.OriginRequestHeaders
			// 		headers[name] = value
			// 		paths = append(paths, endpoint.ConfigPath{
			// 			Path: path,
			// 			Url: configpath.Url,
			// 			OriginRequestHeaders: headers,
			// 		})
			// 	} else {
			// 		paths = append(paths, configpath)
			// 	}
			// }

			// if err := configSrv.Write(configjson); err != nil {
			// 	log.Fatalf("Error: %s", err.Error())
			// }
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