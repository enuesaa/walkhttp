package command

import (
	"fmt"
	"slices"
	"strings"

	"github.com/enuesaa/walkhttp/internal/invoke"
	"github.com/enuesaa/walkhttp/internal/repository"
)

type Prompter struct {
	repos repository.Repos
}

func (p *Prompter) Run() error {
	for {
		entry, err := p.buildreq()
		if err != nil {
			return err
		}
		if err := p.confirm(); err != nil {
			return err
		}
		if err := p.invoke(&entry); err != nil {
			return err
		}
	}
}

func (p *Prompter) buildreq() (invoke.Entry, error) {
	invokeSrv := invoke.New(p.repos)
	entry := invokeSrv.NewEntry("GET", "")

	methods := []string{"get", "post", "put", "delete", "options", "GET", "POST", "PUT", "DELETE", "OPTIONS"}
	validate := func (value string) error {
		if !slices.Contains(methods, strings.ToUpper(value)) {
			return fmt.Errorf("method `%s` is invalid", entry.Request.Method)
		}
		return nil
	}
	if err := p.repos.Prompt.AskSuggest("", "", methods, &entry.Request.Method, validate); err != nil {
		return entry, err
	}
	entry.Request.Method = strings.ToUpper(entry.Request.Method)
	p.repos.Log.Printf("\n")
	p.repos.Log.Printf("│ %s\n", entry.Request.Method)

	if err := p.repos.Prompt.Ask("Url", "", &entry.Request.Url); err != nil {
		return entry, err
	}
	p.repos.Log.Printf("│ %s\n", entry.Request.Url)

	for {
		var headerName string
		var headerValue string
		suggestion := []string{"content-type", "accept"}
		if err := p.repos.Prompt.AskSuggest("Header Name", "(To skip, click enter)", suggestion, &headerName, nil); err != nil {
			return entry, err
		}
		if headerName == "" {
			break
		}

		if err := p.repos.Prompt.Ask("Header Value", fmt.Sprintf(" (%s) ", headerName), &headerValue); err != nil {
			return entry, err
		}
		entry.Request.Headers[headerName] = []string{headerValue}
		p.repos.Log.Printf("│ %s: %s\n", headerName, headerValue)
	}

	if entry.Request.Method == "POST" || entry.Request.Method == "PUT" {
		if err := p.repos.Prompt.Text("Body", "", &entry.Request.Body); err != nil {
			return entry, err
		}
		p.repos.Log.Printf("│\n")
		p.repos.Log.Printf("│ [Body]\n")
		p.repos.Log.Printf("│ %s\n", entry.Request.Body)
	}
	return entry, nil
}

func (p *Prompter) confirm() error {
	confirm := false
	if err := p.repos.Prompt.Confirm("Do you confirm?", &confirm); err != nil {
		return err
	}
	if !confirm {
		return fmt.Errorf("canceled")
	}
	return nil
}

func (p *Prompter) invoke(entry *invoke.Entry) error {
	invokeSrv := invoke.New(p.repos)

	if err := invokeSrv.Invoke(entry); err != nil {
		return err
	}
	p.repos.Log.Printf("│\n")
	p.repos.Log.Printf("│ Status: %d\n", entry.Response.Status)

	return invokeSrv.Save(*entry)
}
