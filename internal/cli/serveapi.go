package cli

import (
	"github.com/enuesaa/walkin/internal/service"
	"github.com/spf13/cobra"
)

func CreateServeApiCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve-api",
		Short: "serve api",
		Run: func(cmd *cobra.Command, args []string) {	
			port, _ := cmd.Flags().GetInt("port")
	
			webSrv := service.NewWebService()
			webSrv.EnableApi()
			webSrv.SetPort(port)
			webSrv.Serve()
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}