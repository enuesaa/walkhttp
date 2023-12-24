package pages

import (
	"encoding/json"
	"path/filepath"
)

type Page struct {
	Filename string // todo
	Method string `json:"method"`
	Url string `json:"url"`
	// request
	// response
}

func (srv *PagesSrv) ReadPage(filename string) (Page, error) {
	var page Page
	path := filepath.Join(srv.workdir, "pages", filename + ".json")
	fbytes, err := srv.repos.Fs.Read(path)
	if err != nil {
		return page, err
	}
	if err := json.Unmarshal(fbytes, &page); err != nil {
		return page, err
	}
	page.Filename = filename
	return page, nil
}
