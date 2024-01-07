package cli

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateRunCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run",
		Run: func(cmd *cobra.Command, args []string) {
			filename, _ := cmd.Flags().GetString("file")

			taskfile, err := usecase.ParseTaskfile(repos, filename)
			if err != nil {
				repos.Log.Fatalf("Error: %s", err.Error())
			}
			if err := usecase.Invoke(repos, taskfile); err != nil {
				repos.Log.Fatalf("Error: %s", err.Error())
			}

			if err := usecase.FlushLog(repos, taskfile); err != nil {
				repos.Log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().String("file", "task.hcl", "taskfile name")

	return cmd
}
