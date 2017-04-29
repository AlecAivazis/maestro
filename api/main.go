package main

import (
	"github.com/graphql-go/graphql"
	"github.com/nautilus/events"
	"github.com/nautilus/events/kafka"
	n "github.com/nautilus/services"
	nHttp "github.com/nautilus/services/http"
)

// MaestroAPI is the service that acts as the interface between the
// various clients and the backend services.
type MaestroAPI struct {
	events.EventBroker
}

// Schema returns the graphql schema associated with this service.
func (s *MaestroAPI) Schema() *graphql.Schema {
	return Schema
}

func main() {
	// try to connect to kafka
	broker, err := KafkaBroker.New(&KafkaBroker.NewOptions{
		Topic: "build",
	})
	if err != nil {
		panic(err)
	}

	// an instance of the service
	service := &MaestroAPI{
		EventBroker: broker,
	}

	// start the event listener
	go n.Start(service, nil)
	// start the api service
	nHttp.Start(service, nil)
}
