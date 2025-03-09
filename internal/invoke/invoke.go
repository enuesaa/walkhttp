package invoke

import (
	"net/http"
	"time"
)

func (srv *InvokeSrv) Invoke(entry *Entry) error {
	client := http.Client{}

	req, err := srv.buildReq(entry)
	if err != nil {
		return err
	}

	// do
	entry.Request.Started = time.Now().UnixMilli()
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	entry.Response.Received = time.Now().UnixMilli()
	entry.HttpVersion = req.Proto

	if err := srv.parseRes(entry, res); err != nil {
		return err
	}

	if srv.repos.Env.WALKHTTP_REQUEST_BODY == "exclude" {
		entry.Request.Body = ""
	}
	if srv.repos.Env.WALKHTTP_RESPONSE_BODY == "exclude" {
		entry.Response.Body = ""
	}

	return nil
}
