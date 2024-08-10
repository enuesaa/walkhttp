package invoke

import (
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

func (srv *InvokeSrv) Save(invocation workspace.Entry) error {
	path := srv.repos.Ws.GetPath()
	ws, err := srv.repos.Ws.Read(path)
	if err != nil {
		return err
	}
	ws.Entries = append(ws.Entries, invocation)

	return srv.repos.Ws.Write(path, ws)
}
