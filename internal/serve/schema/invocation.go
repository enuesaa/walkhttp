package schema

import (
	"fmt"

	"github.com/enuesaa/walkhttp/internal/invoke"
)

func NewInvocationFromEntry(entry invoke.Entry) Invocation {
	invocation := Invocation{
		ID: entry.Id,
		Status: entry.Response.Status,
		Method: entry.Request.Method,
		URL: entry.Request.Url,
		RequestHeaders: make([]*Header, 0),
		ResponseHeaders: make([]*Header, 0),
		RequestBody: entry.Request.Body,
		ResponseBody: entry.Response.Body,
		Created: fmt.Sprint(entry.Request.Started),
	}

	return invocation
}

