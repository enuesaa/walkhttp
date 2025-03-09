package schema

import (
	"cmp"
	"slices"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
)

func NewInvocationFromEntry(entry invoke.Entry) Invocation {
	created := time.UnixMilli(entry.Request.Started).Format(time.RFC3339Nano)
	received := time.UnixMilli(entry.Response.Received).Format(time.RFC3339Nano)

	invocation := Invocation{
		ID:              entry.Id,
		Status:          entry.Response.Status,
		Method:          entry.Request.Method,
		URL:             entry.Request.Url,
		RequestHeaders:  []*Header{},
		ResponseHeaders: []*Header{},
		RequestBody:     entry.Request.Body,
		ResponseBody:    entry.Response.Body,
		Started:         created,
		Received:        received,
	}

	for name, values := range entry.Request.Headers {
		invocation.RequestHeaders = append(invocation.RequestHeaders, &Header{
			Name:  name,
			Value: values,
		})
	}
	for name, values := range entry.Response.Headers {
		invocation.ResponseHeaders = append(invocation.ResponseHeaders, &Header{
			Name:  name,
			Value: values,
		})
	}

	slices.SortFunc(invocation.RequestHeaders, func(a *Header, b *Header) int {
		return cmp.Compare(a.Name, b.Name)
	})
	slices.SortFunc(invocation.ResponseHeaders, func(a *Header, b *Header) int {
		return cmp.Compare(a.Name, b.Name)
	})

	return invocation
}
