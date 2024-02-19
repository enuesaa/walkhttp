package usecase

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/enuesaa/walkin/pkg/repository"
)

func SaveInvocation(repos repository.Repos, invocation invoke.Invocation) error {
	data, err := json.Marshal(invocation)
	if err != nil {
		return err
	}
	u, err := url.Parse(invocation.Url)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("walkin-%s-%s.json", 
		invocation.Method,
		strings.ReplaceAll(strings.ReplaceAll(u.Host, ".", "-"), "/", "-"),
	)
	return repos.Fs.Create(filename, data)
}
