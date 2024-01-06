package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
	"github.com/k0kubun/pp/v3"
)

func CreateListCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list",
		Run: func(cmd *cobra.Command, args []string) {
			m := map[string]string{"foo": "bar", "hello": "world"}
			pp.Print(m)
		},
	}

	return cmd
}

