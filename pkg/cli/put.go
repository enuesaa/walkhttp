package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreatePutCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put",
		Short: "make a put request",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			verbose, _ := cmd.Flags().GetBool("verbose")

			invocation, err := usecase.PromptReq(repos, "PUT", url)
			if err != nil {
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

	return cmd
}