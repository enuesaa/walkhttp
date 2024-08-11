package invoke

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (srv *InvokeSrv) Invoke(entry *Entry) error {
	entry.Request.Started = time.Now().Unix()

	reqbody := bytes.NewBuffer([]byte(entry.Request.Body))
	req, err := http.NewRequest(entry.Request.Method, entry.Request.Url, reqbody)
	if err != nil {
		return err
	}

	for name, values := range entry.Request.Headers {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	entry.Response.Received = time.Now().Unix()
	entry.Response.Status = res.StatusCode
	for key, values := range res.Header {
		if len(values) == 0 {
			return fmt.Errorf("failed to map response header because there is no value supplied")
		}
		entry.Response.Headers[key] = values
	}

	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	entry.Response.Body = string(resbody)

	return nil
}
