package invoke

import (
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

func (srv *InvokeSrv) Read() (workspace.Workspace, error) {
	path := srv.repos.Ws.GetPath()
	return srv.repos.Ws.Read(path)
}
