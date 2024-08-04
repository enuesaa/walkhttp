package usecase

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

func LoadConfig(repos repository.Repos, path string) workspace.Workspace {
	if path == "" {
		return workspace.New()
	}
	ws, err := repos.Fs.ReadWorkspace(path)
	if err != nil {
		return workspace.New()
	}
	return ws
}

func WriteConfig(repos repository.Repos, path string, ws workspace.Workspace) error {
	return repos.Fs.WriteWorkspace(path, ws)
}
