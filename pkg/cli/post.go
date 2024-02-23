package cli

import (
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreatePostCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post",
		Short: "make a post request",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			verbose, _ := cmd.Flags().GetBool("verbose")
			browser, _ := cmd.Flags().GetBool("browser")

			if browser {
				// TODO set url, method here.
				return usecase.Serve()
			}

			invocation := invoke.NewInvocation("POST", url)
			if err := usecase.PromptReq(repos, &invocation); err != nil {
				return err
			}
			if err := usecase.Invoke(repos, &invocation, !verbose); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("url", "", "url")
	cmd.Flags().BoolP("verbose", "v", false, "verbose")
	cmd.Flags().Bool("browser", false, "serve web console to request in a browser")

	return cmd
}
