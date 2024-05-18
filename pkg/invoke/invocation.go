package invoke

import (
	"github.com/oklog/ulid/v2"
)

func (srv *InvokeSrv) Create(method string, url string) Invocation {
	return Invocation{
		ID:              ulid.Make().String(),
		Method:          method,
		URL:             url,
		RequestHeaders:  []*Header{},
		ResponseHeaders: []*Header{},
	}
}
