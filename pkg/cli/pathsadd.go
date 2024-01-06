package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreatePathsAddCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")
			url, _ := cmd.Flags().GetString("url")

			fmt.Printf("%s: %s\n", path, url)
		},
	}
	cmd.Flags().String("path", "/", "proxy path")
	cmd.Flags().String("url", "https://example.com/", "proxy url")

	cmd.MarkFlagRequired("path")
	cmd.MarkFlagRequired("url")

	return cmd
}

