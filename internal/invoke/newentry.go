package invoke

import "fmt"

func (srv *InvokeSrv) NewEntry(method string, path string) Entry {
	url := fmt.Sprintf("%s%s", srv.BaseUrl(), path)

	data := Entry{
		Id: "",
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
