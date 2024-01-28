package usecase

import (
	"fmt"

	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/gofiber/fiber/v2"
)

type CreateInvokeResponseData struct {}

type CreateInvokeRequest struct {
	Method string `json:"method"`
	Url string `json:"url"`
	Body string `json:"body"`
}
func (ctl *ServeCtl) CreateInvoke(c *fiber.Ctx) error {
	req := CreateInvokeRequest{}
    if err := c.BodyParser(&req); err != nil {
        return err
    }

	invocation := invoke.Invocation {
		Method: req.Method,
		Url: req.Url,
		RequestBody: []byte(req.Body),
	}
	if err := invoke.Invoke(&invocation); err != nil {
		return err
	}
	fmt.Printf("invoke: %d\n", invocation.Status)

	return c.JSON(ServeEmptyResponse{})
}
