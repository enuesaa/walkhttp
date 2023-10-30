package cli

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/enuesaa/walkin/internal/service"
)

// see https://pkg.go.dev/github.com/graph-gophers/graphql-go@v1.5.0/relay#Handler.ServeHTTP
// { "query": "{ fileinfo(name: \"aa\") { name, description } }" }
type QueryRequestSchema struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func CreateServeGraphqlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve-graphql",
		Short: "serve graphql",
		Run: func(cmd *cobra.Command, args []string) {	
			port, _ := cmd.Flags().GetInt("port")

			app := fiber.New()
			app.Post("/query", func(c *fiber.Ctx) error {
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
			})

			addr := fmt.Sprintf(":%d", port)
			app.Listen(addr)
		},
	}
	cmd.Flags().Int("port", 3000, "port")

	return cmd
}
