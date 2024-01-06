package cli

import (
	"log"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateInvokeCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			if err := usecase.Invoke(repos, name); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().String("name", "", "name")
	cmd.MarkFlagRequired("name")

	return cmd
}

