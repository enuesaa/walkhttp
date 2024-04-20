package serve

import (

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/labstack/echo/v4"
)

func (s *Servectl) handleApi(c echo.Context) error {
	invocation := invoke.NewInvocation("GET", "https://example.com")
	if err := invoke.Invoke(&invocation); err != nil {
		return err
	}
	return nil
}
