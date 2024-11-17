package invoke

import (
	"fmt"
	"io"
	"net/http"
)

func (srv *InvokeSrv) parseRes(entry *Entry, res *http.Response) error {
	// status
	entry.Response.Status = res.StatusCode

	// headers
	for key, values := range res.Header {
		if len(values) == 0 {
			return fmt.Errorf("failed to map response header because there is no value supplied")
		}
		entry.Response.Headers[key] = values
	}

	// body
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	entry.Response.Body = string(resbody)

	return nil
}
