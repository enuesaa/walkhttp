package invoke

import (
	"net/http"
	"strings"
)

func (srv *InvokeSrv) buildReq(entry *Entry) (*http.Request, error) {
	reqbody := strings.NewReader(entry.Request.Body)

	req, err := http.NewRequest(entry.Request.Method, entry.Request.Url, reqbody)
	if err != nil {
		return req, err
	}

	for name, values := range entry.Request.Headers {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	return req, nil
}
