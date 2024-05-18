package invoke

import (
	"encoding/json"
)

func (srv *InvokeSrv) GetLog(id string) (*Invocation, error) {
	path, err := srv.GetLogFilename(id)
	if err != nil {
		return nil, err
	}
	fbyte, err := srv.repos.Fs.Read(path)
	if err != nil {
		return nil, err
	}
	var invocation Invocation
	if err := json.Unmarshal(fbyte, &invocation); err != nil {
		return nil, err
	}
	return &invocation, nil
}
