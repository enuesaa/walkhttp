package schema

import (
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
)

func NewInvocationFromEntry(entry invoke.Entry) Invocation {
	created := time.Unix(entry.Request.Started, 0).Format(time.RFC3339)

	invocation := Invocation{
		ID:              entry.Id,
		Status:          entry.Response.Status,
		Method:          entry.Request.Method,
		URL:             entry.Request.Url,
		RequestHeaders:  make([]*Header, 0),
		ResponseHeaders: make([]*Header, 0),
		RequestBody:     entry.Request.Body,
		ResponseBody:    entry.Response.Body,
		Created:         created,
	}

	for name, values := range entry.Request.Headers {
		entry.Request.Headers[name] = values
	}
	for name, values := range entry.Response.Headers {
		entry.Response.Headers[name] = values
	}

	return invocation
}
