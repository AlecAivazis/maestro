package main

import (
	"net/http"

	"github.com/nautilus/services/graphql"
	goji "goji.io"
)

// the routes that this service
func (s *MaestroAPI) Router() http.Handler {
	// create an empty mux we can play with
	mux := goji.NewMux()

	// add the graphl routes to it
	GraphqlService.AddRoutes(s, mux)

	// return the router we just made
	return mux
}
