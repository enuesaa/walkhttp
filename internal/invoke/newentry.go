package invoke

func (srv *InvokeSrv) NewEntry(method string) Entry {
	url := srv.GetBaseUrl()

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
