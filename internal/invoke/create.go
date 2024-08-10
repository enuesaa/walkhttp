package invoke

import (
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
	"github.com/google/uuid"
)

func (srv *InvokeSrv) Create(method string, url string) workspace.Entry {
	return workspace.Entry{
		Id: uuid.NewString(),
		Request: workspace.EntryRequest{},
		Response: workspace.EntryResponse{},
	}
}
