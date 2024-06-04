package invoke

import (
	"encoding/json"
)

func (srv *InvokeSrv) CreateLog(invocation *Invocation) error {
	path := srv.GetLogFilename(invocation.ID)

	data, err := json.Marshal(invocation)
	if err != nil {
		return err
	}
	return srv.repos.Fs.Create(path, data)
}
