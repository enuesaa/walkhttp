package schema

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
)

func (i *BrowserInvocationInput) ToEntry() invoke.Entry {
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
			Status: i.Status,
			Headers: make(map[string][]string),
			Body: i.ResponseBody,
			Received: 0,    
		},
	}

	return entry
}
