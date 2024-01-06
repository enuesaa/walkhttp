package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePathsAddOriginRequestHeaderCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-origin-request-header",
		Short: "add header of origin request",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			name, _ := cmd.Flags().GetString("name")
			value, _ := cmd.Flags().GetString("value")

			fmt.Printf("%s: %s, %s\n", path, name, value)
		},
	}
	cmd.Flags().String("path", "/", "proxy path")
	cmd.Flags().String("name", "Accept", "header name")
	cmd.Flags().String("value", "application/json", "header value")

	cmd.MarkFlagRequired("path")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("value")

	return cmd
}