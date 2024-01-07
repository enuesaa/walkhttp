package invoke

type ResponseHeader struct {
	Key   string
	Value string
}

type Response struct {
	Status  int
	Headers []ResponseHeader
	Body    []byte
}
