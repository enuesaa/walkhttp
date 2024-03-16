package usecase

import (
	"github.com/enuesaa/walkin/pkg/repository"
	"github.com/enuesaa/walkin/pkg/serve"
)

func Serve(repos repository.Repos, port int) error {
	servectl := serve.New(repos)
	servectl.Port = port

	return servectl.Listen()
}
