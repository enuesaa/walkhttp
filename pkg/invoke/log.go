package invoke

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/enuesaa/walkin/pkg/repository"
)

func CreateLog(repos repository.Repos, invocation *Invocation) error {
	homedir, err := repos.Fs.HomeDir()
	if err != nil {
		return err
	}
	walkindir := filepath.Join(homedir, ".walkin")
	logsdir := filepath.Join(walkindir, "logs")

	if err := repos.Fs.CreateDir(logsdir); err != nil {
		return err
	}
	path := filepath.Join(logsdir, fmt.Sprintf("%s.json", invocation.ID))

	data, err := json.Marshal(invocation)
	if err != nil {
		return err
	}

	return repos.Fs.Create(path, data)
}
