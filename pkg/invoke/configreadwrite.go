package invoke

import (
	"encoding/json"
)

func (srv *InvokeSrv) Read(path string) (Config, error) {
	fbytes, err := srv.repos.Fs.Read(path)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := json.Unmarshal(fbytes, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (srv *InvokeSrv) Write(path string, config Config) error {
	fbytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err := srv.repos.Fs.Create(path, fbytes); err != nil {
		return  err
	}
	return nil
}
