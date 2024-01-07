package usecase

import (
	"fmt"
	"time"

	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/taskparse"
)

func FlushLog(repos repository.Repos, taskfile taskparse.Taskfile) error {
	logs := repos.Log.GetAll()

	now := time.Now()
	logFilename := fmt.Sprintf("walkin-%s-%s.log", taskfile.Batch.Name, now.Format("20060102150405"))
	if err := repos.Fs.Create(logFilename, []byte(logs)); err != nil {
		return fmt.Errorf("Error: failed to create log file")
	}

	return nil
}
