package cli

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateOptionsCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "options <path>",
		Short: "make a options request",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := ""
			if len(args) > 0 {
				path = args[0]
			}

			invocation := invoke.NewInvocation("OPTIONS", path)
			if err := usecase.PromptReq(repos, &invocation); err != nil {
				return err
			}
			if err := usecase.Invoke(repos, &invocation, false); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
