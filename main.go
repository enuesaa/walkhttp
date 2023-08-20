package main

import (
	// "flag"
	// "fmt"

	"fmt"

	"github.com/gofiber/fiber/v2"
	graphql "github.com/graph-gophers/graphql-go"
)

type query struct{}

func (_ *query) Hey() string { return "Hey, world!" }

func main() {
	// flag.Parse()
	// args := flag.Args()
	// if len(args) == 0 {
	// 	fmt.Println("Required argument missing: WORKDIR")
	// 	return
	// }

	// workdir := args[0]
	// fmt.Println(workdir)

	app := fiber.New()

	// _ := app.Group("/api")
	// create api routes here

	// see https://pkg.go.dev/github.com/graph-gophers/graphql-go@v1.5.0/relay#Handler.ServeHTTP
	type QueryRequestSchema struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	app.Post("/query", func(c *fiber.Ctx) error {
		body := new(QueryRequestSchema)
		if err := c.BodyParser(body); err != nil {
			return err
		}

		fmt.Printf("%+v", body)

		// definition
		s := `
		type Query {
			hey: String!
		}
		`
		schema := graphql.MustParseSchema(s, &query{})

		response := schema.Exec(
			c.Context(),
			body.Query,
			body.OperationName,
			body.Variables,
		)
	
		return c.JSON(response)
	})
	// http.Handle("/query", &relay.Handler{Schema: schema})
	// http.ListenAndServe(":3000", nil)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	// should return html or assets
	// 	return c.SendString("")
	// })
	app.Listen(":3000")
}
