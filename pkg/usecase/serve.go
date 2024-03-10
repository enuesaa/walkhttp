package usecase

import (
	"github.com/enuesaa/walkin/pkg/servectl"
)

func Serve(port int) error {
	server := servectl.Servectl{
		Wsconns: servectl.NewWsConns(),
	}

	return server.Listen(port)
}
