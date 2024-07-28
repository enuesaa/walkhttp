package invoke

import (
	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

func (srv *InvokeSrv) Save(invocation schema.Invocation) error {
	return srv.repos.DB.Save(invocation.ID, invocation)
}
