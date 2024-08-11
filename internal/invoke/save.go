package invoke

func (srv *InvokeSrv) Save(entry Entry) error {
	if !srv.repos.Fs.IsExist(srv.repos.WorkspacePath) {
		ws := Workspace{
			BaseUrl: "",
		}
		if err := srv.Write(ws); err != nil {
			return err
		}
	}

	ws, err := srv.Read()
	if err != nil {
		return err
	}
	ws.Entries = append(ws.Entries, entry)

	return srv.Write(ws)
}
