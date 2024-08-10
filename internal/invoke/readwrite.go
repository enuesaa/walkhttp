package invoke

import (
	"encoding/json"
)

func (srv *InvokeSrv) Read() (Workspace, error) {
	var ws Workspace
	
	fbytes, err := srv.repos.Fs.Read(srv.repos.WorkspacePath)
	if err != nil {
		return ws, err
	}

	if err := json.Unmarshal(fbytes, &ws); err != nil {
		return ws, err
	}
	return ws, nil
}

func (srv *InvokeSrv) Write(ws Workspace) error {
	fbytes, err := json.MarshalIndent(ws, "", "  ")
	if err != nil {
		return err
	}
	return srv.repos.Fs.Create(srv.repos.WorkspacePath, fbytes)
}

func (srv *InvokeSrv) Save(entry Entry) error {
	ws, err := srv.Read()
	if err != nil {
		return err
	}
	ws.Entries = append(ws.Entries, entry)

	return srv.Write(ws)
}
