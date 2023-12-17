package event

type Event struct {
	Url string // include queries
	Request EventRequest
	Response EventResponse
}

type EventRequest struct {
	Headers map[string]string
	Body string
}

type EventResponse struct {
	Status int
	Headers map[string]string
	Body string
}
