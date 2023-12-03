package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/internal/repository"
	"github.com/spf13/cobra"
)

func CreateMockCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mock",
		Short: "mock",
		Run: func(cmd *cobra.Command, args []string) {	
			portFlag, _ := cmd.Flags().GetInt("port")
			fmt.Printf("%d\n", portFlag)

			// read mock files.

			fmt.Printf("mocking endpoints below.\n")
			// list endpoints
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}
