package invoke

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (srv *InvokeSrv) Invoke(invocation *Entry) error {
	invocation.Id = uuid.NewString()
	invocation.Request.Started = time.Now().Unix()

	reqbody := bytes.NewBuffer([]byte(invocation.Request.Body))
	req, err := http.NewRequest(invocation.Request.Method, invocation.Request.Url, reqbody)
	if err != nil {
		return err
	}

	for name, values := range invocation.Request.Headers {
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

	invocation.Response.Status = res.StatusCode
	for key, values := range res.Header {
		if len(values) == 0 {
			return fmt.Errorf("failed to map response header because there is no value supplied")
		}
		invocation.Response.Headers[key] = values
	}

	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	invocation.Response.Body = string(resbody)

	return nil
}
