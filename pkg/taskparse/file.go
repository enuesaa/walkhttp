package taskparse

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Taskfile struct {
	Batch Batch `hcl:"batch,block"`
}
type Batch struct {
	Name    string `hcl:"name"`
	Comment string `hcl:"comment"`
	Tasks   []Task `hcl:"task,block"`
}
type Task struct {
	Name       string       `hcl:"name,label"`
	TaskRun    []TaskRun    `hcl:"run,block"`
	TaskInvoke []TaskInvoke `hcl:"invoke,block"`
}
type TaskRun struct {
	Command []string `hcl:"command"`
}
type TaskInvoke struct {
	Url string `hcl:"url"`
}

func (srv *TaskParser) Read(filename string) (Taskfile, error) {
	var taskfile Taskfile
	fbytes, err := srv.repos.Fs.Read(filename)
	if err != nil {
		return taskfile, err
	}
	if err := hclsimple.Decode(filename, fbytes, nil, &taskfile); err != nil {
		return taskfile, err
	}
	return taskfile, nil
}

func (srv *TaskParser) IsConfigFileExist() bool {
	return srv.repos.Fs.IsExist("walkin.json")
}
