package cli

import (
	"github.com/enuesaa/walkin/internal/repository"
	"github.com/enuesaa/walkin/internal/web"
	"github.com/spf13/cobra"
)

func CreateUpCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "up web server",
		Run: func(cmd *cobra.Command, args []string) {	
			portFlag, _ := cmd.Flags().GetInt("port")
			proxyFlag, _ := cmd.Flags().GetString("proxy")
			readLocalFilesFlag, _ := cmd.Flags().GetString("read-local-files")
			configFlag, _ := cmd.Flags().GetString("config")

			serveConfig := web.ServeConfig{}
			if configFlag != "" {
			} else {
				if proxyFlag != "" {
					serveConfig.Paths["/*"] = web.Behavior{
						Behavior: web.Proxy,
						ProxyConfig: web.ProxyConfig{
							Url: "https://example.com/",
						},
					}
				}
				if readLocalFilesFlag != "" {
					serveConfig.Paths["/*"] = web.Behavior{
						Behavior: web.ReadLocalFiles,
					}
				}
			}
	
			webSrv := web.NewWebService(repos)
			webSrv.WithPort(portFlag)
			webSrv.WithServeConfig(serveConfig)
			webSrv.Serve()
		},
	}
	cmd.Flags().Int("port", 3000, "port")
	cmd.Flags().String("proxy", "path=/*,url=https://example.com", "serve reverse proxy on specific path")
	cmd.Flags().String("read-local-files", "", "serve local files")
	cmd.Flags().String("config", "", "config file path")

	return cmd
}
