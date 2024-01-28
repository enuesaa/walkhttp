package invoke

type Invocation struct {
	Status int `json:"status"`
	Method string `json:"method"`
	Url string `json:"url"`

	RequestHeaders []Header `json:"requestHeaders"`
	RequestBody []byte `json:"requestBody"`

	ResponseBody []byte `json:"responseBody"`
	ResponseHeaders []Header `json:"responseHeaders"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
