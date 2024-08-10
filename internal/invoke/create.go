package invoke

import (
	"github.com/google/uuid"
)

func (srv *InvokeSrv) Create(method string, url string) Entry {
	data := Entry{
		Id: uuid.NewString(),
		Request: EntryRequest{
			Method: method,
			Url: url,
			Headers: map[string][]string{},
			Body: "",
			Started: 0,
		},
		Response: EntryResponse{
			Status: 0,
			Headers: map[string][]string{},
			Body: "",
			Received: 0,
		},
	}

	return data
}
