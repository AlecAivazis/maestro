package main

import (
	"net/http"

	"github.com/nautilus/services/graphql"
)

// the routes that this service
func (s *MaestroAPI) Router() http.Handler {
	// create an empty mux we can play with
	mux := http.NewServeMux()

	// add the graphl routes to the service mux
	GraphqlService.AddRoutes(s, mux)

	// return the router we just made
	return mux
}
