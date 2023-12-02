package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/internal/repository"
	"github.com/enuesaa/walkin/internal/web"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve web server",
		Run: func(cmd *cobra.Command, args []string) {	
			portFlag, _ := cmd.Flags().GetInt("port")
			proxyFlag, _ := cmd.Flags().GetStringArray("proxy")
			localFilesFlag, _ := cmd.Flags().GetStringArray("local-files")

			serveConfig, err := ParseFlagsToServeConfig(localFilesFlag, proxyFlag)
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			webSrv := web.NewWebService(repos)
			webSrv.WithPort(portFlag)
			webSrv.WithServeConfig(serveConfig)
			webSrv.Serve()
		},
	}
	cmd.Flags().Int("port", 3000, "port")
	cmd.Flags().StringArray("proxy", []string{}, "serve reverse proxy on specific path like `path=/*,url=https://example.com`")
	cmd.Flags().StringArray("local-files", []string{}, "serve local files like `path=/*`")

	return cmd
}
