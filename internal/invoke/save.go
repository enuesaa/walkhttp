package invoke

func (srv *InvokeSrv) Save(entry Entry) error {
	return srv.repos.DB.Write(entry.Id, entry)
}
