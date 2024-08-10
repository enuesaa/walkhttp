package invoke

func (srv *InvokeSrv) Get(id string) (Entry, error) {
	ws, err := srv.Read()
	if err != nil {
		return Entry{}, err
	}

	for _, entry := range ws.Entries {
		if entry.Id == id {
			return entry, nil
		}
	}

	return Entry{}, nil
}
