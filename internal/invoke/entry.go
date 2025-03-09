package invoke

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type Entry struct {
	Id          string        `json:"id"`
	Request     EntryRequest  `json:"request"`
	Response    EntryResponse `json:"response"`
	HttpVersion string        `json:"httpVersion"`
}

type EntryRequest struct {
	Method  string              `json:"method"`
	Url     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
	Started int64               `json:"started"`
}

type EntryResponse struct {
	Status   int                 `json:"status"`
	Headers  map[string][]string `json:"headers"`
	Body     string              `json:"body"`
	Received int64               `json:"received"`
}

func (srv *InvokeSrv) NewEntry(method string, path string) Entry {
	url := fmt.Sprintf("%s%s", srv.BaseUrl(), path)

	data := Entry{
		Id: ulid.Make().String(),
		Request: EntryRequest{
			Method:  method,
			Url:     url,
			Headers: map[string][]string{},
			Body:    "",
			Started: 0,
		},
		Response: EntryResponse{
			Status:   0,
			Headers:  map[string][]string{},
			Body:     "",
			Received: 0,
		},
	}
	return data
}
