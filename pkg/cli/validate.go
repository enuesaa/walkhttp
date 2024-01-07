package cli

import (
	"log"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
)

type TaskRunCommand struct {
	Name string      `hcl:"name,label"`
	Command []string `hcl:"command"`
}

type Task struct {
    Name    string `hcl:"name"`
    Comment string `hcl:"comment"`
	TaskRunCommands []TaskRunCommand `hcl:"run_command,block"`
}

type Taskfile struct {
    Task Task `hcl:"task,block"`
}

func CreateValidateCmd(repos repository.Repos) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "validate config files.",
		Run: func(cmd *cobra.Command, args []string) {
			var taskConfig Taskfile
			err := hclsimple.DecodeFile("testdata/normal/task1.hcl", nil, &taskConfig)
			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			pp.Print(taskConfig)
		},
	}

	return cmd
}

