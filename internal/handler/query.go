package handler

import (
	"fmt"

	"github.com/enuesaa/walkin/internal/service"
	"github.com/gofiber/fiber/v2"
)

// see https://pkg.go.dev/github.com/graph-gophers/graphql-go@v1.5.0/relay#Handler.ServeHTTP
// { "query": "{ fileinfo(name: \"aa\") { name, description } }" }
type QueryRequestSchema struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

type QueryController struct{}

func (ctl *QueryController) Query(c *fiber.Ctx) error {
	body := new(QueryRequestSchema)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	response, err := service.ExecQuery(body.Query, body.OperationName, body.Variables)
	if err != nil {
		fmt.Println(err)
		return c.JSON("{}")
	}

	return c.JSON(response)
}
