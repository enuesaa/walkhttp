package cli

import (
	"github.com/enuesaa/walkin/internal/service"
	"github.com/spf13/cobra"
)

func CreateUpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "up web server",
		Run: func(cmd *cobra.Command, args []string) {	
			port, _ := cmd.Flags().GetInt("port")
	
			webSrv := service.NewWebService()
			webSrv.EnableStaticServe()
			webSrv.SetPort(port)
			webSrv.Serve()
		},
	}
	cmd.Flags().Int("port", 3000, "port")
	cmd.Flags().StringArray("proxy", []string{}, "serve reverse proxy on specific path")
	cmd.Flags().StringArray("read-local-files", []string{}, "serve local files.")

	return cmd
}
