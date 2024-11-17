package invoke

import (
	"io"
	"net/http"
	"unicode/utf8"
)

func (srv *InvokeSrv) parseRes(entry *Entry, res *http.Response) error {
	// status
	entry.Response.Status = res.StatusCode

	// headers
	for key, values := range res.Header {
		if len(values) == 0 {
			continue
		}
		entry.Response.Headers[key] = values
	}

	// body
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if utf8.Valid(resbody) {
		entry.Response.Body = string(resbody)
	}

	return nil
}
