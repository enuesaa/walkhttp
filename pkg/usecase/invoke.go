package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/taskparse"
)

func Invoke(repos repository.Repos, taskfile taskparse.Taskfile) error {
	invokeSrv := invoke.NewInvokeSrv(repos)
	for _, task := range taskfile.Batch.Tasks {
		if len(task.TaskInvoke) == 0 {
			continue
		}
		invokeConfig := task.TaskInvoke[0]
		request := invoke.Request{
			Method:  "GET",
			Url:     invokeConfig.Url,
			Headers: make([]invoke.RequestHeader, 0),
			Body:    make([]byte, 0),
		}
		response, err := invokeSrv.Invoke(request)
		if err != nil {
			return fmt.Errorf("failed to invoke because `%s`", err.Error())
		}

		repos.Log.Printf("status: %d\n", response.Status)
		repos.Log.Printf("body: %s\n", response.Body)
	}

	return nil
}
