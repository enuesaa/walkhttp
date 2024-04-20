package usecase

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func Invoke(repos repository.Repos, invocation *invoke.Invocation, truncate bool) error {
	if err := invoke.Invoke(invocation); err != nil {
		return err
	}

	data, err := json.Marshal(invocation)
	if err != nil {
		return err
	}
	operationName, err := invocation.GetOperationName()
	if err != nil {
		return err
	}
	fmt.Printf("status: %d\n", invocation.Status)

	homedir, err := repos.Fs.HomeDir()
	if err != nil {
		return err
	}
	walkindir := filepath.Join(homedir, ".walkin")
	logsdir := filepath.Join(walkindir, "logs")

	if err := repos.Fs.CreateDir(logsdir); err != nil {
		return err
	}
	filename := fmt.Sprintf("%s.json", operationName)
	path := filepath.Join(logsdir, filename)

	if err := repos.Fs.Create(path, data); err != nil {
		return err
	}

	return nil
}
