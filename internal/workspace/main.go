package workspace

import (
	"encoding/json"

	"github.com/enuesaa/walkhttp/internal/repository"
)

func NewSrv(repos repository.Repos) WorkspaceSrv {
	return WorkspaceSrv{
		repos: repos,
	}
}

type WorkspaceSrv struct {
	repos repository.Repos
}

func (srv *WorkspaceSrv) Read(path string) (Workspace, error) {
	fbytes, err := srv.repos.Fs.Read(path)
	if err != nil {
		return Workspace{}, err
	}
	var config Workspace
	if err := json.Unmarshal(fbytes, &config); err != nil {
		return Workspace{}, err
	}
	return config, nil
}

func (srv *WorkspaceSrv) Write(path string, config Workspace) error {
	fbytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err := srv.repos.Fs.Create(path, fbytes); err != nil {
		return err
	}
	return nil
}
