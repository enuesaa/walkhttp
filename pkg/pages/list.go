package pages

import (
	"fmt"
	"path/filepath"
	"strings"
)

func (srv *PagesSrv) ListPageFilenames() ([]string, error) {
	var list []string

	pagesDir := filepath.Join(srv.workdir, "pages")
	filenames, err := srv.repos.Fs.ListFiles(pagesDir)
	if err != nil {
		return list, fmt.Errorf("Error: failed to list files")
	}

	// trim workdir path
	for _, filename := range filenames {
		rel, err := filepath.Rel(pagesDir, filename)
		if err != nil {
			return list, err
		}
		path := strings.TrimSuffix(rel, ".json")
		list = append(list, path)
	}

	return list, err
}

func (srv *PagesSrv) ListPages() ([]Page, error) {
	list := make([]Page, 0)

	filenames, err := srv.ListPageFilenames()
	if err != nil {
		return list, err
	}

	for _, filename := range filenames {
		page, err := srv.ReadPage(filename)
		if err != nil {
			return list, err
		}
		list = append(list, page)
	}

	return list, err
}
