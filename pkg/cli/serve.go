package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve",
		Run: func(cmd *cobra.Command, args []string) {	
			// serve fiber app here.

			// on request,
			// read config file, and get origin url.
			// request to origin
			// save response (as history and cache)

			// queue mutation request.
		},
	}

	return cmd
}
