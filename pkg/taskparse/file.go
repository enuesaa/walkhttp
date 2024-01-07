package taskparse

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Taskfile struct {
    Task Task `hcl:"task,block"`
}
type Task struct {
    Name    string `hcl:"name"`
    Comment string `hcl:"comment"`
	TaskRunCommands []TaskRunCommand `hcl:"run_command,block"`
	TaskInvoke []TaskInvoke `hcl:"invoke,block"`
}
type TaskRunCommand struct {
	Name string      `hcl:"name,label"`
	Command []string `hcl:"command"`
}
type TaskInvoke struct {
	Name string `hcl:"name,label"`
	Url string `hcl:"url"`
}

func (srv *TaskParser) Read() (Taskfile, error) {
	var taskfile Taskfile
	fbytes, err := srv.repos.Fs.Read("testdata/normal/task1.hcl")
	if err != nil {
		return taskfile, err
	}
	if err := hclsimple.Decode("task1.hcl", fbytes, nil, &taskfile); err != nil {
		return taskfile, err
	}
	return taskfile, nil
}

func (srv *TaskParser) IsConfigFileExist() bool {
	return srv.repos.Fs.IsExist("walkin.json")
}
