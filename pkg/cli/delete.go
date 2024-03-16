package cli

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateDeleteCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <path>",
		Short: "make a delete request",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CheckConfigFileExists(repos)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			path := ""
			if len(args) > 0 {
				path = args[0]
			}

			invocation := invoke.NewInvocation("DELETE", path)
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
