package schema

func NewConfig() Config {
	return Config{
		BaseUrl: "https://",
	}
}

type Config struct {
	BaseUrl string `json:"baseUrl"` // like `https://example.com`
	// Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Path string `json:"path"` // like `/aaa`
	Method string `json:"method"` // like `GET`
	Headers []Header `json:"headers"` // default headers
	Body string `json:"body"` // default body
}