package schema

import (
	"cmp"
	"slices"
	"time"

	"github.com/enuesaa/walkhttp/internal/invoke"
)

func NewInvocationFromEntry(entry invoke.Entry) Invocation {
	created := time.Unix(entry.Request.Started, 0).Format(time.RFC3339)
	received := time.Unix(entry.Response.Received, 0).Format(time.RFC3339)

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
		for _, value := range values {
			invocation.RequestHeaders = append(invocation.RequestHeaders, &Header{
				Name:  name,
				Value: value,
			})
		}
	}
	for name, values := range entry.Response.Headers {
		for _, value := range values {
			invocation.ResponseHeaders = append(invocation.ResponseHeaders, &Header{
				Name:  name,
				Value: value,
			})
		}
	}

	slices.SortFunc(invocation.RequestHeaders, func(a *Header, b *Header) int {
		return cmp.Compare(a.Name, b.Name)
	})
	slices.SortFunc(invocation.ResponseHeaders, func(a *Header, b *Header) int {
		return cmp.Compare(a.Name, b.Name)
	})

	return invocation
}
