package main

import (
	"github.com/nautilus/events"
	"github.com/nautilus/services"
)

// MaestroRepo is the service responsible for retrieving information for a
// given repo
type MaestroRepo struct {
	events.EventBroker
}

func main() {

	// try to connect to kafka
	broker, err := events.NewKafkaBroker(&events.NewKafkaBrokerOptions{
		Topic: "repo",
	})
	if err != nil {
		panic(err)
	}

	// for now, just run a single build instance
	service := MaestroRepo{
		EventBroker: broker,
	}

	// start the event listener
	Service.Start(&service, &Service.RuntimeConfig{
		EventBroker: broker,
	})
}
