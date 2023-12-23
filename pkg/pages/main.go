package pages

import "github.com/enuesaa/walkin/pkg/repository"

func NewPagesSrv(repos repository.Repos, workdir string) PagesSrv {
	return PagesSrv{
		repos: repos,
		workdir: workdir,
	}
}

type PagesSrv struct {
	repos repository.Repos
	workdir string
}
