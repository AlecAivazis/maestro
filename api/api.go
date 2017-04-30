package main

import (
	"net/http"

	"github.com/alecaivazis/maestro/api/mutations"
	"github.com/alecaivazis/maestro/api/objectTypes"
	"github.com/graphql-go/graphql"
	"github.com/nautilus/events"
	"github.com/nautilus/services/graphql"
)

// MaestroAPI is the service that acts as the interface between the
// various clients and the backend services.
type MaestroAPI struct {
	events.EventBroker
}

// Schema is the graphql schema for the api provided by this service.
var Schema *graphql.Schema

// the schema definition
var schemaConfig = graphql.SchemaConfig{
	Query:    objectTypes.Query,
	Mutation: mutations.Mutations,
}

// Schema returns the graphql schema associated with this service.
func (s *MaestroAPI) Schema() *graphql.Schema {
	return Schema
}

// the routes that this service
func (s *MaestroAPI) Router() http.Handler {
	// create an empty mux we can play with
	mux := http.NewServeMux()

	// add the graphl routes to the service mux
	GraphqlService.AddRoutes(s, mux)

	// return the router we just made
	return mux
}

// when this package is loaded
func init() {
	// define the schema
	schema, err := graphql.NewSchema(schemaConfig)
	// if something went wrong
	if err != nil {
		panic(err)
	}

	Schema = &schema
}
