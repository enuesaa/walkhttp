package invoke

import (
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

func (srv *InvokeSrv) Save(invocation Invocation) error {
	path := srv.repos.Ws.GetPath()
	ws, err := srv.repos.Ws.Read(path)
	if err != nil {
		return err
	}
	ws.Entries = append(ws.Entries, workspace.Entry{
		Request: workspace.EntryRequest{
			Method: invocation.Method,
			Url: invocation.URL,
			Headers: map[string][]string{},
			Body: invocation.RequestBody,
			Started: 0,
		},
		Response: workspace.EntryResponse{
			Status: invocation.Status,
			Body: invocation.ResponseBody,
		},
	})

	return srv.repos.Ws.Write(path, ws)
}
