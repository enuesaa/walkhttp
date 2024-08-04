package invoke

import (
	"github.com/oklog/ulid/v2"
)

//Deprecated
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//Deprecated
type Invocation struct {
	ID              string    `json:"id"`
	Status          int       `json:"status"`
	Method          string    `json:"method"`
	URL             string    `json:"url"`
	RequestHeaders  []*Header `json:"requestHeaders,omitempty"`
	ResponseHeaders []*Header `json:"responseHeaders,omitempty"`
	RequestBody     string    `json:"requestBody"`
	ResponseBody    string    `json:"responseBody"`
	Created         string    `json:"created"`
}

//Deprecated
func (srv *InvokeSrv) Create(method string, url string) Invocation {
	return Invocation{
		ID:              ulid.Make().String(),
		Method:          method,
		URL:             url,
		RequestHeaders:  []*Header{},
		ResponseHeaders: []*Header{},
	}
}
