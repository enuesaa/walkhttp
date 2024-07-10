package invoke

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/enuesaa/walkhttp/pkg/serve/schema"
)

func (srv *InvokeSrv) Invoke(invocation *schema.Invocation) error {
	invocation.Created = time.Now().Format(time.RFC3339)

	reqbody := bytes.NewBuffer([]byte(invocation.RequestBody))
	req, err := http.NewRequest(invocation.Method, invocation.URL, reqbody)
	if err != nil {
		return err
	}

	for _, requestHeader := range invocation.RequestHeaders {
		req.Header.Add(requestHeader.Name, requestHeader.Value)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	invocation.Status = res.StatusCode
	for key, value := range res.Header {
		if len(value) == 0 {
			return fmt.Errorf("failed to map response header because there is no value supplied.")
		}
		invocation.ResponseHeaders = append(invocation.ResponseHeaders, &schema.Header{
			Name:  key,
			Value: value[0],
		})
	}

	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	invocation.ResponseBody = string(resbody)

	return nil
}
