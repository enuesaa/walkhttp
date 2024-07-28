package invoke

func (srv *InvokeSrv) Save(invocation Invocation) error {
	return srv.repos.DB.Save(invocation.ID, invocation)
}
