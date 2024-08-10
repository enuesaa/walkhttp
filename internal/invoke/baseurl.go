package invoke

func (srv *InvokeSrv) GetBaseUrl() string {
	ws, err := srv.Read()
	if err != nil {
		return "https://example.com/"
	}
	return ws.BaseUrl
}
