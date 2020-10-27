package main

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func main() {
	// setup GraphQL schema and resolver
	s := `
		type Query {
						hello: String!
		}
	`
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})

	// register graphql playground handler
	http.Handle("/", http.FileServer(http.Dir("./playground")))

	// start web server
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
