package invoke

import (
	"encoding/json"
)

func (srv *InvokeSrv) ListLogs() ([]*Invocation, error) {
	list := make([]*Invocation, 0)

	logsdir, err := srv.GetLogsDir()
	if err != nil {
		return list, err
	}
	files, err := srv.repos.Fs.ListFiles(logsdir)
	if err != nil {
		return list, err
	}
	for _, path := range files {
		fbyte, err := srv.repos.Fs.Read(path)
		if err != nil {
			return list, err
		}
		var invocation Invocation
		if err := json.Unmarshal(fbyte, &invocation); err != nil {
			return list, err
		}
		list = append(list, &invocation)
	}

	return list, nil
}
