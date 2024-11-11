package invoke

func (srv *InvokeSrv) Get(id string) (Entry, error) {
	data, err := srv.repos.DB.Read(id)
	if err != nil {
		return Entry{}, err
	}
	return data.(Entry), nil
}
