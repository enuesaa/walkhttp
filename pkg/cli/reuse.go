package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateReuseCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reuse",
		Short: "reuse old request",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			verbose, _ := cmd.Flags().GetBool("verbose")

			filename := args[0]
			fmt.Printf("file: %s\n", filename)

			// reuse old invocations

			// mock
			invocation := invoke.Invocation {
				Method: "GET",
				Url: "https://example.com",
				RequestHeaders: make([]invoke.Header, 0),
				ResponseHeaders: make([]invoke.Header, 0),
			}
			if err := usecase.Invoke(repos, &invocation, !verbose); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().BoolP("verbose", "v", false, "verbose")

	return cmd
}
