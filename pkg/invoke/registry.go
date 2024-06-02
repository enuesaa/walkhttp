package invoke

import (
	"fmt"
	"path/filepath"
)

func (srv *InvokeSrv) GetLogsDir() string {
	return "."
}

func (srv *InvokeSrv) CreateLogsDir() error {
	logsdir := srv.GetLogsDir()
	return srv.repos.Fs.CreateDir(logsdir)
}

func (srv *InvokeSrv) GetLogFilename(id string) (string, error) {
	logsdir := srv.GetLogsDir()
	filename := fmt.Sprintf("%s.json", id)
	path := filepath.Join(logsdir, filename)
	return path, nil
}
