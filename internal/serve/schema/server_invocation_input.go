package schema

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
)

func (i *ServerInvocationInput) ToEntry() invoke.Entry {
	entry := invoke.Entry{
		Id: "",
		Request: invoke.EntryRequest{
			Method: i.Method,
			Url: i.URL,
			Headers: make(map[string][]string),
			Body: i.RequestBody,
			Started: 0, // TODO
		},
		Response: invoke.EntryResponse{
			Status: 0,
			Headers: make(map[string][]string),
			Body: "",
			Received: 0,    
		},
	}

	return entry
}
