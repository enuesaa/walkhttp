package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/enuesaa/walkin/pkg/endpoint"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func Invoke(repos repository.Repos, name string) error {
	endpointSrv := endpoint.New(repos)
	configfile, err := endpointSrv.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file")
	}

	endpointdef := endpoint.Endpoint{
		Name: "",
	}
	for _, def := range configfile.Endpoints {
		if def.Name == name {
			endpointdef = def
			break
		}
	}
	if endpointdef.Name == "" {
		return fmt.Errorf("failed to find endpoint with name %s", name)
	}
	fmt.Printf("%+v\n", endpointdef)

	invokeSrv := invoke.NewInvokeSrv(repos)

	request := invoke.Request{
		Method: "GET",
		Url: endpointdef.Url,
		Headers: make([]invoke.RequestHeader, 0),
		Body: make([]byte, 0),
	}
	for headerName, headerValue := range endpointdef.RequestHeaders {
		request.Headers = append(request.Headers, invoke.RequestHeader{
			Key: headerName,
			Value: headerValue,
		})
	}

	reqbodybytes, err := json.Marshal(endpointdef.RequestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body")
	}
	request.Body = reqbodybytes
	response, err := invokeSrv.Invoke(request)
	if err != nil {
		return fmt.Errorf("failed to invoke because `%s`", err.Error())
	}
	fmt.Printf("status: %d\n", response.Status)
	fmt.Printf("body: %s\n", response.Body)

	return nil
}