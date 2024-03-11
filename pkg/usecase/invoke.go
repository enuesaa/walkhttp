package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func Invoke(repos repository.Repos, invocation *invoke.Invocation, truncate bool) error {
	if err := invoke.Invoke(invocation); err != nil {
		return err
	}

	data, err := json.Marshal(invocation)
	if err != nil {
		return err
	}
	operationName, err := invocation.GetOperationName()
	if err != nil {
		return err
	}
	fmt.Printf("status: %d\n", invocation.Status)

	// TODO
	// if truncate is true, truncate response body with 100 letters.
	if !truncate {
		fmt.Printf("body: %s\n", invocation.ResponseBody)
	}

	filename := fmt.Sprintf("walkin-%s.json", operationName)
	if err := repos.Fs.Create(filename, data); err != nil {
		return err
	}

	return nil
}
