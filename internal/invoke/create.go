package invoke

import (
	"github.com/enuesaa/walkhttp/internal/repository/workspace"
	"github.com/google/uuid"
)

func (srv *InvokeSrv) Create(method string, url string) workspace.Entry {
	data := workspace.Entry{
		Id: uuid.NewString(),
		Request: workspace.EntryRequest{
			Method: method,
			Url: url,
			Headers: map[string][]string{},
			Body: "",
			Started: 0,
		},
		Response: workspace.EntryResponse{
			Status: 0,
			Headers: map[string][]string{},
			Body: "",
			Received: 0,
		},
	}

	return data
}
