package usecase

import (
	"github.com/enuesaa/walkin/pkg/pages"
	"github.com/enuesaa/walkin/pkg/repository"
)

func ListPages(repos repository.Repos, workdir string) ([]pages.Page, error) {
	list := make([]pages.Page, 0)
	pagesSrv := pages.NewPagesSrv(repos)
	pagesSrv.SetWorkdir(workdir)

	filenames, err := pagesSrv.ListPageFilenames()
	if err != nil {
		return list, err
	}

	for _, filename := range filenames {
		page, err := pagesSrv.ReadPage(filename)
		if err != nil {
			return list, err
		}
		list = append(list, page)
	}

	return list, err
}