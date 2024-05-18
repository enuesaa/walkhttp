package invoke

import (
	"fmt"
	"path/filepath"
)

func (srv *InvokeSrv) GetWalkinDir() (string, error) {
	homedir, err := srv.repos.Fs.HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".walkin"), nil
}

func (srv *InvokeSrv) GetLogsDir() (string, error) {
	walkindir, err := srv.GetWalkinDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(walkindir, "logs"), nil
}

func (srv *InvokeSrv) CreateLogsDir() error {
	logsdir, err := srv.GetLogsDir()
	if err != nil {
		return err
	}
	return srv.repos.Fs.CreateDir(logsdir)
}

func (srv *InvokeSrv) GetLogFilename(id string) (string, error) {
	logsdir, err := srv.GetLogsDir()
	if err != nil {
		return "", err
	}
	filename := fmt.Sprintf("%s.json", id)
	path := filepath.Join(logsdir, filename)
	return path, nil
}
