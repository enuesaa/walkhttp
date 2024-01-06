package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/endpoint"
	"github.com/enuesaa/walkin/pkg/repository"
)

func CreateConfigFileIfNotExist(repos repository.Repos) error {
	endpointSrv := endpoint.New(repos)
	if endpointSrv.IsConfigFileExist() {
		return nil
	}

	configfile := endpoint.ConfigFile{
		Endpoints: make(map[string]endpoint.Endpoint, 0),
	}
	return endpointSrv.WriteConfig(configfile)
}

func DefineEndpointWithPrompt(repos repository.Repos) (endpoint.Endpoint, error) {
	def := endpoint.Endpoint{
		RequestHeaders: make(map[string]string, 0),
		RequestBody: make(map[string]string, 0),
	}
	
	url, err := repos.Prompt.Ask("url: ", "https://example.com/")
	if err != nil {
		return def, err
	}
	def.Url = url

	for {
		choice, err := repos.Prompt.Select("", []string{"header", "body", "end"})
		if err != nil {
			break
		}

		switch choice {
		case "header":
			headerName, err := repos.Prompt.Ask("headerName: ", "Content-Type")
			if err != nil {
				return def, err
			}
			msg := fmt.Sprintf("headerValue(%s): ", headerName)
			headerValue, err := repos.Prompt.Ask(msg, "application/json")
			if err != nil {
				return def, err
			}
			def.RequestHeaders[headerName] = headerValue
		case "body":
			bodyPath, err := repos.Prompt.Ask("bodyPath: ", "$.aa")
			if err != nil {
				return def, err
			}
			msg := fmt.Sprintf("bodyValue(%s): ", bodyPath)
			bodyValue, err := repos.Prompt.Ask(msg, "aa")
			if err != nil {
				return def, err
			}
			def.RequestBody[bodyPath] = bodyValue
		}

		if choice == "end" {
			break
		}
	}

	return def, nil
}