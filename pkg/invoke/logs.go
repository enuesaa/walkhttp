package invoke

import (
	"encoding/json"
	"path/filepath"

	"github.com/enuesaa/walkin/pkg/repository"
)

func ListLogs(repos repository.Repos) ([]Invocation, error) {
	list := make([]Invocation, 0)

	homedir, err := repos.Fs.HomeDir()
	if err != nil {
		return list, err
	}
	walkindir := filepath.Join(homedir, ".walkin")
	logsdir := filepath.Join(walkindir, "logs")

	files, err := repos.Fs.ListFiles(logsdir)
	if err != nil {
		return list, err
	}
	for _, path := range files {
		fbyte, err := repos.Fs.Read(path)
		if err != nil {
			return list, err
		}
		var invocation Invocation
		if err := json.Unmarshal(fbyte, &invocation); err != nil {
			return list, err
		}
		list = append(list, invocation)
	}

	return list, nil
}
