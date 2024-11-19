package invoke

func (srv *InvokeSrv) Lifecycle() error {
	overCount := srv.repos.DB.Count() - srv.repos.Env.WALKHTTP_PERSIST_COUNT
	if overCount <= 0 {
		return nil
	}

	list, err := srv.List()
	if err != nil {
		return err
	}

	for i := range overCount {
		data := list[len(list)-i-1]
		if err := srv.Delete(data.Id); err != nil {
			return err
		}
		srv.repos.Log.Printf("Lifecycle: data deleted `%s`\n", data.Id)
	}
	return nil
}
