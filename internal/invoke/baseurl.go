package invoke

import "os"

func (srv *InvokeSrv) BaseUrl() string {
	url := os.Getenv("WALKHTTP_URL_BASE")

	if url == "" {
		// default value
		return "https://example.com/"
	}
	return url
}
