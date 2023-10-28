package cli

import (
	"github.com/spf13/cobra"
	"github.com/enuesaa/walkin/internal/handler"
)

func CreateServeGraphqlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve-graphql",
		Short: "serve graphql",
		Run: func(cmd *cobra.Command, args []string) {	
			port, _ := cmd.Flags().GetInt("port")
			handler.Handle(port)
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}