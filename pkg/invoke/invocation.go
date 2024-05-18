package invoke

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// todo: make usecase
func NewInvocation(method string, url string) Invocation {
	return Invocation{
		Method:          method,
		Url:             url,
		RequestHeaders:  []Header{},
		ResponseHeaders: []Header{},
	}
}

type Invocation struct {
	Status int    `json:"status"`
	Method string `json:"method"`
	Url    string `json:"url"`

	RequestHeaders []Header `json:"requestHeaders"`
	RequestBody    string   `json:"requestBody"`

	ResponseBody    string   `json:"responseBody"`
	ResponseHeaders []Header `json:"responseHeaders"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (inv *Invocation) GetOperationName() (string, error) {
	u, err := url.Parse(inv.Url)
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
