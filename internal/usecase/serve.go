package usecase

import (
	"github.com/enuesaa/walkhttp/internal/repository"
	"github.com/enuesaa/walkhttp/internal/serve"
)

func Serve(repos repository.Repos, port int) error {
	serveCtl := serve.New(repos)
	serveCtl.UsePort(port)

	return serveCtl.Serve()
}
