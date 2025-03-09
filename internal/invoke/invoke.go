package invoke

import (
	"fmt"
	"net/http"
	"time"
)

func (srv *InvokeSrv) Invoke(entry *Entry) error {
	entry.Request.Started = time.Now().UnixMilli()

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
	entry.Response.Received = time.Now().UnixMilli()
	fmt.Printf("%+v\n", entry)
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
