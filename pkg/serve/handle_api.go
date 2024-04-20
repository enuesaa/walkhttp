package serve

import (

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/gofiber/fiber/v2"
)

func (s *Servectl) handleApi(c *fiber.Ctx) error {
	// todo
	url := c.OriginalURL()

	invocation := invoke.NewInvocation(c.Method(), url)
	if err := invoke.Invoke(&invocation); err != nil {
		return err
	}
	go func() {
		s.wsconns.Send(invocation.String())
	}()

	return nil
}
