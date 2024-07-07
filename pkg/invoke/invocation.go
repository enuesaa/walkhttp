package invoke

import (
	"github.com/enuesaa/walkhttp/pkg/schema"
	"github.com/oklog/ulid/v2"
)

func (srv *InvokeSrv) Create(method string, url string) schema.Invocation {
	return schema.Invocation{
		ID:              ulid.Make().String(),
		Method:          method,
		URL:             url,
		RequestHeaders:  []*schema.Header{},
		ResponseHeaders: []*schema.Header{},
	}
}
