package pages

import "github.com/enuesaa/walkin/pkg/repository"

func NewPagesSrv(repos repository.Repos) PagesSrv {
	return PagesSrv{
		repos: repos,
		workdir: ".",
	}
}

type PagesSrv struct {
	repos repository.Repos
	workdir string
}

func (srv *PagesSrv) SetWorkdir(workdir string) {
	srv.workdir = workdir
}
