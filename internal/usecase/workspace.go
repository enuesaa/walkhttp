package usecase

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

func ReadWorkspace(repos repository.Repos, path string) workspace.Workspace {
	if path == "" {
		return workspace.Workspace{
			BaseUrl: "https://example.com/",
			Entries: make([]workspace.Entry, 0),
		}
	}
	ws, err := repos.Ws.Read(path)
	if err != nil {
		return workspace.Workspace{
			BaseUrl: "https://example.com/",
			Entries: make([]workspace.Entry, 0),
		}
	}
	return ws
}

func WriteWorkspace(repos repository.Repos, path string, ws workspace.Workspace) error {
	return repos.Ws.Write(path, ws)
}
