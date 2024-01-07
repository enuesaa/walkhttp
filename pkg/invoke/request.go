package invoke

type RequestHeader struct {
	Key   string
	Value string
}

type Request struct {
	Method  string
	Url     string
	Headers []RequestHeader
	Body    []byte
}
