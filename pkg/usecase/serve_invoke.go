package usecase

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/enuesaa/walkin/pkg/invoke"
	"github.com/gofiber/fiber/v2"
)

type CreateInvokeResponseData struct {}

type CreateInvokeRequest struct {
	Method string `json:"method" validate:"required,oneof=GET POST PUT DELETE OPTIONS"`
	Url string `json:"url" validate:"required,url"`
	Body string `json:"body"`
}
func (ctl *ServeCtl) CreateInvoke(c *fiber.Ctx) error {
	req := CreateInvokeRequest{}
    if err := c.BodyParser(&req); err != nil {
        return err
    }
	if err := validator.New().Struct(req); err != nil {
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
