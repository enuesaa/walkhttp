package schema

import (
	"github.com/enuesaa/walkhttp/internal/invoke"
)

func (i *ServerInvocationInput) ToEntry() invoke.Entry {
	entry := invoke.Entry{
		Id: "", // TODO
		Request: invoke.EntryRequest{
			Method:  i.Method,
			Url:     i.URL,
			Headers: map[string][]string{},
			Body:    i.RequestBody,
			Started: 0,
		},
		Response: invoke.EntryResponse{
			Status:   0,
			Headers:  map[string][]string{},
			Body:     "",
			Received: 0,
		},
	}

	for _, h := range i.RequestHeaders {
		if _, ok := entry.Request.Headers[h.Name]; !ok {
			entry.Request.Headers[h.Name] = []string{}
		}
		entry.Request.Headers[h.Name] = append(entry.Request.Headers[h.Name], h.Value)
	}

	return entry
}
