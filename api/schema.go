package main

import (
	"github.com/graphql-go/graphql"

	"github.com/alecaivazis/maestro/api/mutations"
	"github.com/alecaivazis/maestro/api/objectTypes"
)

// Schema is the graphql schema for the api provided by this service.
var Schema *graphql.Schema

// the schema definition
var schemaConfig = graphql.SchemaConfig{
	Query:    objectTypes.Query,
	Mutation: mutations.Mutations,
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
