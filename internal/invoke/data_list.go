package invoke

import "sort"

func (srv *InvokeSrv) List() ([]Entry, error) {
	list := []Entry{}

	for _, key := range srv.repos.DB.Keys() {
		data, err := srv.repos.DB.Read(key)
		if err != nil {
			return list, err
		}
		list = append(list, data.(Entry))
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Request.Started > list[j].Request.Started
	})

	return list, nil
}
