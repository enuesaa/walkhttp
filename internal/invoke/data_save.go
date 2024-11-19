package invoke

func (srv *InvokeSrv) Save(entry Entry) error {
	if err := srv.repos.DB.Write(entry.Id, entry); err != nil {
		return err
	}
	return srv.Lifecycle()
}
