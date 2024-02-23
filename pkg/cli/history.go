package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateHistoryCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "history",
		Short: "lookup histories",
		RunE: func(cmd *cobra.Command, args []string) error {
			filenames, err := repos.Fs.ListFiles(".")
			if err != nil {
				return err
			}
			
			for _, filename := range filenames {
				if !strings.HasPrefix(filename, "walkin-") {
					continue
				}

				data, err := repos.Fs.Read(filename)
				if err != nil {
					continue
				}
				var invocation invoke.Invocation
				if err := json.Unmarshal(data, &invocation); err != nil {
					return err
				}
				fmt.Printf("%d %s %s\n", invocation.Status, invocation.Method, invocation.Url)
			}

			return nil
		},
	}

	return cmd
}