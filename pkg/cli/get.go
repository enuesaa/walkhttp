package cli

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateGetCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "make a get request",
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			verbose, _ := cmd.Flags().GetBool("verbose")

			invocation, err := usecase.PromptReq(repos, "GET", url)
			if err != nil {
				return err
			}

			if err := usecase.Invoke(repos, &invocation); err != nil {
				return err
			}
			fmt.Printf("status: %d\n", invocation.Status)
			if verbose {
				fmt.Printf("body: %s\n", invocation.ResponseBody)				
			}
			if err := usecase.SaveInvocation(repos, invocation); err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().String("url", "", "url")
	cmd.Flags().BoolP("verbose", "v", false, "verbose")

	return cmd
}
