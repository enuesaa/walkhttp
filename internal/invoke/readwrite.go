package invoke

func (srv *InvokeSrv) Read() (Workspace, error) {
	var ws Workspace

	for _, key := range srv.repos.DB.Keys() {
		data, err := srv.repos.DB.Read(key)
		if err != nil {
			return ws, err
		}
		ws.Entries = append(ws.Entries, data.(Entry))
	}

	return ws, nil
}

func (srv *InvokeSrv) Write(entry Entry) error {
	return srv.repos.DB.Write(entry.Id, entry)
}
