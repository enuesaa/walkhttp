package endpoint

import (
	"encoding/json"
)

type ConfigFile struct {
	Endpoints map[string]Endpoint `json:"endpoints"`
}
type Endpoint struct {
	Url string `json:"url"`
	RequestHeaders map[string]string `json:"requestHeaders"`
}

func (srv *EndpointSrv) GetConfigFilename() string {
	return "walkin.json"
}

func (srv *EndpointSrv) ReadConfig() (ConfigFile, error) {
	var config ConfigFile
	fbytes, err := srv.repos.Fs.Read(srv.GetConfigFilename())
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(fbytes, &config); err != nil {
		return config, err
	}
	return config, nil
}

func (srv *EndpointSrv) IsConfigFileExist() bool {
	return srv.repos.Fs.IsExist(srv.GetConfigFilename())
}

func (srv *EndpointSrv) WriteConfig(config ConfigFile) error {
	fbytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return srv.repos.Fs.Create(srv.GetConfigFilename(), fbytes)
}
