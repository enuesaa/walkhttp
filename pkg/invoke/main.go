package invoke

import (
	"io"
	"net/http"

	"github.com/enuesaa/walkin/pkg/repository"
)

func NewInvokeSrv(repos repository.Repos) InvokeSrv {
	return InvokeSrv{
		repos: repos,
	}
}

type InvokeSrv struct {
	repos repository.Repos
}

func (srv *InvokeSrv) Invoke(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return make([]byte, 0), err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}
	defer res.Body.Close()
	resbodybytes, err := io.ReadAll(res.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	return resbodybytes, err
}
