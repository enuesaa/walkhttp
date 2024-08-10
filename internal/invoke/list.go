package invoke

//TODO: paginate
func (srv *InvokeSrv) List() ([]Entry, error) {
	ws, err := srv.Read()
	if err != nil {
		return []Entry{}, err
	}

	return ws.Entries, nil
}
