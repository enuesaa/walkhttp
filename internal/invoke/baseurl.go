package invoke

func (srv *InvokeSrv) BaseUrl() string {
	return srv.repos.Env.WALKHTTP_URL_BASE
}
