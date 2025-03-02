package prompt

import (
	"fmt"
	"slices"
	"strings"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

func BuildReq(repos repository.Repos, entry *invoke.Entry) error {
	methods := []string{"get", "post", "put", "delete", "options", "GET", "POST", "PUT", "DELETE", "OPTIONS"}
	validate := func (value string) error {
		if !slices.Contains(methods, strings.ToUpper(value)) {
			return fmt.Errorf("method `%s` is not defined", entry.Request.Method)
		}
		return nil
	}
	if err := repos.Prompt.AskSuggest("", "", methods, &entry.Request.Method, validate); err != nil {
		return err
	}
	entry.Request.Method = strings.ToUpper(entry.Request.Method)

	repos.Log.Printf("│ %s\n", entry.Request.Method)
	if err := repos.Prompt.Ask("Url", "", &entry.Request.Url); err != nil {
		return err
	}
	repos.Log.Printf("│ %s\n", entry.Request.Url)
	repos.Log.Printf("│\n")

	for {
		var headerName string
		var headerValue string
		suggestion := []string{"content-type", "accept"}
		if err := repos.Prompt.AskSuggest("Header Name", "(To skip, click enter)", suggestion, &headerName, nil); err != nil {
			return err
		}
		if headerName == "" {
			break
		}

		if err := repos.Prompt.Ask("Header Value", fmt.Sprintf(" (%s) ", headerName), &headerValue); err != nil {
			return err
		}
		entry.Request.Headers[headerName] = []string{headerValue}
		repos.Log.Printf("│ %s: %s\n", headerName, headerValue)
	}

	if entry.Request.Method == "POST" || entry.Request.Method == "PUT" {
		if err := repos.Prompt.Text("Body", "", &entry.Request.Body); err != nil {
			return err
		}
		repos.Log.Printf("│\n")
		repos.Log.Printf("│ [Body]\n")
		repos.Log.Printf("│ %s\n", entry.Request.Body)
	}
	return nil
}
