package cli

import (
	"github.com/enuesaa/walkin/internal/web"
	"github.com/spf13/cobra"
)

func CreateUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "up web server",
		Run: func(cmd *cobra.Command, args []string) {	
			port, _ := cmd.Flags().GetInt("port")
	
			webSrv := web.NewWebService()
			webSrv.WithPort(port)
			webSrv.Serve()
		},
	}
	cmd.Flags().Int("port", 3000, "port")
	cmd.Flags().String("proxy", "path=/*,url=https://example.com", "serve reverse proxy on specific path")
	cmd.Flags().String("read-local-files", "", "serve local files")

	return cmd
}
