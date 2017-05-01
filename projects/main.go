package main

import (
	"github.com/nautilus/events"
	"github.com/nautilus/services"
)

func main() {
	// try to connect to kafka
	broker, err := events.NewKafkaBroker(&events.NewKafkaBrokerOptions{
		Topic: "projects",
	})
	if err != nil {
		panic(err)
	}

	// for now, just run a single build instance
	service := MaestroProjects{
		EventBroker: broker,
	}

	// start the event listener
	Service.Start(&service, &Service.RuntimeConfig{
		EventBroker: broker,
	})
}
