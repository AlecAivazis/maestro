package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	graphiql "github.com/mnmtanish/go-graphiql"
)

func main() {
	// attach the graphql endpoints
	http.Handle("/graphql", handler.New(&handler.Config{
		Schema: Schema,
		Pretty: true,
	}))
	http.HandleFunc("/graphiql", graphiql.ServeGraphiQL)
	// start the server
	log.Fatal(http.ListenAndServe(":4000", nil))
}
