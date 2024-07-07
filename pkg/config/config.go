package config

import (
	"github.com/enuesaa/walkhttp/pkg/invoke"
)

type Config struct {
	BaseUrl string `json:"baseUrl"` // like `https://example.com`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Path string `json:"path"` // like `/aaa`
	Method string `json:"method"` // like `GET`
	Headers []invoke.Header `json:"headers"` // default headers
	Body string `json:"body"` // default body
}
