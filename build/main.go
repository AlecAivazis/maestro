package main

import (
	"github.com/nautilus/events"
	"github.com/nautilus/events/kafka"
	"github.com/nautilus/services"
)

// MaestroBuild is the service responsible for building projects, running any
// test scripts, and reporting the status as well as any logs along the way.
type MaestroBuild struct {
	events.EventBroker
}

func main() {

	// try to connect to kafka
	broker, err := KafkaBroker.New(&KafkaBroker.NewOptions{
		Topic: "build",
	})
	if err != nil {
		panic(err)
	}

	// for now, just run a single build instance
	service := MaestroBuild{
		EventBroker: broker,
	}

	// start the event listener
	Service.Start(&service, &Service.RuntimeConfig{
		EventBroker: broker,
	})
}
