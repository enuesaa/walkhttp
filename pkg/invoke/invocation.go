package invoke

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/oklog/ulid/v2"
)

func NewInvocation(method string, url string) Invocation {
	return Invocation{
		ID:              ulid.Make().String(),
		Method:          method,
		URL:             url,
		RequestHeaders:  []*Header{},
		ResponseHeaders: []*Header{},
	}
}

func (inv *Invocation) GetOperationName() (string, error) {
	u, err := url.Parse(inv.URL)
	if err != nil {
		return "", err
	}
	hyphenedHost := strings.ReplaceAll(u.Host, ".", "-")
	hyphenedPath := strings.ReplaceAll(strings.ReplaceAll(u.Path, ".", "-"), "/", "-")

	return fmt.Sprintf("%s-%s-%s", inv.Method, hyphenedHost, hyphenedPath), nil
}

func (inv *Invocation) String() string {
	data, err := json.Marshal(inv)
	if err != nil {
		return "{}"
	}
	return string(data)
}
