package usecase

import (
	"github.com/enuesaa/walkin/pkg/serve"
)

func Serve(port int) error {
	servectl := serve.New()
	servectl.Port = port

	return servectl.Listen()
}
