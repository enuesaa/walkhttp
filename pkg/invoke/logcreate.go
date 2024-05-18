package invoke

import (
	"encoding/json"
)

func (srv *InvokeSrv) CreateLog(invocation *Invocation) error {
	if err := srv.CreateLogsDir(); err != nil {
		return err
	}
	path, err := srv.GetLogFilename(invocation.ID)
	if err != nil {
		return err
	}

	data, err := json.Marshal(invocation)
	if err != nil {
		return err
	}
	return srv.repos.Fs.Create(path, data)
}
