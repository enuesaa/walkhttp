package invoke

import (
	"fmt"
	"net/url"
	"strings"
)

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

func NewInvocation(method string, url string) Invocation {
	return Invocation{
		Method:          method,
		Url:             url,
		RequestHeaders:  []Header{},
		ResponseHeaders: []Header{},
	}
}

func (invocation *Invocation) GetOperationName() (string, error) {
	u, err := url.Parse(invocation.Url)
	if err != nil {
		return "", err
	}
	hyphenedHost := strings.ReplaceAll(u.Host, ".", "-")
	hyphenedPath := strings.ReplaceAll(strings.ReplaceAll(u.Path, ".", "-"), "/", "-")

	return fmt.Sprintf("%s-%s-%s", invocation.Method, hyphenedHost, hyphenedPath), nil
}
