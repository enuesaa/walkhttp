package cli

import (
	"github.com/enuesaa/walkin/internal/service"
	"github.com/spf13/cobra"
)

func CreateServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve static files",
		Run: func(cmd *cobra.Command, args []string) {	
			port, _ := cmd.Flags().GetInt("port")
	
			webSrv := service.NewWebService()
			webSrv.EnableStaticServe()
			webSrv.SetPort(port)
			webSrv.Serve()
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}
