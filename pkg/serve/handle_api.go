package serve

import (
	"encoding/json"
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/gofiber/fiber/v2"
)

func (s *Servectl) handleApi(c *fiber.Ctx) error {
	config, err := s.repos.Conf.Read()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s%s", config.BaseUrl, c.OriginalURL())

	invocation := invoke.NewInvocation(c.Method(), url)
	if err := invoke.Invoke(&invocation); err != nil {
		return err
	}
	data, err := json.Marshal(invocation)
	if err != nil {
		return err
	}
	go func() {
		s.wssend <- data
	}()

	return nil
}
