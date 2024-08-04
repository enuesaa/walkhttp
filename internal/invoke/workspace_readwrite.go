package invoke

import (
	"encoding/json"
)

func (srv *InvokeSrv) Read(path string) (Workspace, error) {
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

func (srv *InvokeSrv) Write(path string, config Workspace) error {
	fbytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err := srv.repos.Fs.Create(path, fbytes); err != nil {
		return err
	}
	return nil
}
