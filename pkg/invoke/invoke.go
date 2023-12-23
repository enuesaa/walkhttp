package invoke

import (
	"fmt"
	"io"
	"net/http"
)

func (srv *InvokeSrv) Invoke(request Request) (Response, error) {
	var response Response

	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		return response, err
	}
	for _, requestHeader := range request.Headers {
		req.Header.Add(requestHeader.Key, requestHeader.Value)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	response.Status = res.StatusCode
	for key, value := range res.Header {
		if len(value) == 0 {
			return response, fmt.Errorf("failed to map response header because there is no value supplied.")
		}
		response.Headers = append(response.Headers, ResponseHeader{
			Key: key,
			Value: value[0],
		})
	}

	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	response.Body = resbody

	return response, nil
}
