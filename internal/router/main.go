package router

import (
	"fmt"

	"github.com/enuesaa/walkhttp/internal/repository"
)

func New(repos repository.Repos) ServeCtl {
	return ServeCtl{
		repos: repos,
		port:  3000,
	}
}

type ServeCtl struct {
	repos repository.Repos
	port  int
}

func (ctl *ServeCtl) UsePort(port int) {
	ctl.port = port
}

func (ctl *ServeCtl) Address() string {
	return fmt.Sprintf(":%d", ctl.port)
}
