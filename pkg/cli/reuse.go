package cli

import (
	"encoding/json"
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
			filename := args[0]
			verbose, _ := cmd.Flags().GetBool("verbose")

			data, err := repos.Fs.Read(filename)
			if err != nil {
				return err
			}
			fmt.Printf("found: %s\n", filename)

			var invocation invoke.Invocation
			if err := json.Unmarshal(data, &invocation); err != nil {
				return err
			}

			if err := usecase.PromptReqConfirmOnly(repos, &invocation); err != nil {
				return err
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
