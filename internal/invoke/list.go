package invoke

func (srv *InvokeSrv) List() ([]Entry, error) {
	list := []Entry{}

	for _, key := range srv.repos.DB.Keys() {
		data, err := srv.repos.DB.Read(key)
		if err != nil {
			return list, err
		}
		list = append(list, data.(Entry))
	}
	return list, nil
}

// TODO: paginate
