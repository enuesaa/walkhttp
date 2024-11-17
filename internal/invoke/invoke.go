package invoke

import (
	"net/http"
	"time"
)

func (srv *InvokeSrv) Invoke(entry *Entry) error {
	entry.Request.Started = time.Now().Unix()

	// req
	req, err := srv.buildReq(entry)
	if err != nil {
		return err
	}

	// do
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// res
	entry.Response.Received = time.Now().Unix()
	if err := srv.parseRes(entry, res); err != nil {
		return err
	}

	return nil
}
