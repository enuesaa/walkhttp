package cli

import (
	"github.com/enuesaa/walkin/internal/repository"
	"github.com/spf13/cobra"
)

func CreateMockCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mock",
		Short: "mock",
		Run: func(cmd *cobra.Command, args []string) {	
			// read mock files.
			// list endpoints
		},
	}

	return cmd
}
