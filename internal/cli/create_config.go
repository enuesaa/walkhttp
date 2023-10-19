package cli

import (
	"github.com/spf13/cobra"
)

func CreateCreateConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-config",
		Short: "create config file with prompt.",
		Run: func(cmd *cobra.Command, args []string) {
			// .walkin 配下に保存する
		},
	}

	return cmd
}
