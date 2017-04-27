package main

import (
	"github.com/graphql-go/graphql"
	nautilus "github.com/nautilus/services/http"
)

// MaestroAPI is the service that acts as the interface between the
// various clients and the backend services.
type MaestroAPI struct{}

// Schema returns the graphql schema associated with this service.
func (s *MaestroAPI) Schema() *graphql.Schema {
	return Schema
}

func main() {
	// start the api service
	nautilus.Start(&MaestroAPI{}, nil)
}
