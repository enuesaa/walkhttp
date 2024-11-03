package invoke

import (
	"path/filepath"

	"github.com/google/uuid"
)

func (srv *InvokeSrv) Save(entry Entry) error {
	if !srv.repos.Fs.IsExist(srv.repos.WorkspacePath) {
		basedir := filepath.Dir(srv.repos.WorkspacePath)
		if err := srv.repos.Fs.CreateDir(basedir); err != nil {
			return err
		}
		ws := Workspace{}
		if err := srv.Write(ws); err != nil {
			return err
		}
	}

	entry.Id = uuid.NewString()

	ws, err := srv.Read()
	if err != nil {
		return err
	}
	ws.Entries = append(ws.Entries, entry)

	return srv.Write(ws)
}
