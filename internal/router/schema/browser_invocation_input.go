package schema

import (
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
)

func (i *BrowserInvocationInput) ToEntry() invoke.Entry {
	started, err := time.Parse(time.RFC3339, i.Started)
	if err != nil {
		started = time.Unix(0, 0)
	}
	received, err := time.Parse(time.RFC3339, i.Received)
	if err != nil {
		started = time.Unix(0, 0)
	}

	entry := invoke.Entry{
		Id: "",
		Request: invoke.EntryRequest{
			Method:  i.Method,
			Url:     i.URL,
			Headers: map[string][]string{},
			Body:    i.RequestBody,
			Started: started.Unix(),
		},
		Response: invoke.EntryResponse{
			Status:   i.Status,
			Headers:  map[string][]string{},
			Body:     i.ResponseBody,
			Received: received.Unix(),
		},
	}

	for _, h := range i.RequestHeaders {
		if _, ok := entry.Request.Headers[h.Name]; !ok {
			entry.Request.Headers[h.Name] = []string{}
		}
		entry.Request.Headers[h.Name] = append(entry.Request.Headers[h.Name], h.Value)
	}

	for _, h := range i.ResponseHeaders {
		if _, ok := entry.Response.Headers[h.Name]; !ok {
			entry.Response.Headers[h.Name] = []string{}
		}
		entry.Response.Headers[h.Name] = append(entry.Response.Headers[h.Name], h.Value)
	}

	return entry
}
