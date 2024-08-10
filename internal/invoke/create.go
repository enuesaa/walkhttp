package invoke

import (
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
)

func (srv *InvokeSrv) Create(method string, url string) workspace.Entry {
	return workspace.Entry{
		Request: workspace.EntryRequest{},
		Response: workspace.EntryResponse{},
	}
}
