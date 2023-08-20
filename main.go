package main

import (
	// "flag"
	// "fmt"

	// "github.com/gofiber/fiber/v2"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
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

	// app := fiber.New()

	// _ := app.Group("/api")
	// create api routes here

	s := `
	type Query {
		hey: String!
	}
	`

	schema := graphql.MustParseSchema(s, &query{})

	// app.Post("/query", func(c *fiber.Ctx) error {
	// 	handler := &relay.Handler{Schema: schema}
	// 	handler.ServeHTTP(c.Response(), c.Request())
	// 	c.SendString()
	// 	return nil
	// })
	http.Handle("/query", &relay.Handler{Schema: schema})
	http.ListenAndServe(":3000", nil)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	// should return html or assets
	// 	return c.SendString("")
	// })
	// app.Listen(":3000")
}
