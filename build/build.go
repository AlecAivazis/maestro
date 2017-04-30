package main

import (
	"fmt"

	"github.com/nautilus/events"

	"github.com/AlecAivazis/maestro/common"
)

// MaestroRepo is the service responsible for retrieving information for a
// given repo
type MaestroRepo struct {
	events.EventBroker
}

func (s *MaestroRepo) HandleAction(a *events.Action) {
	// what we do with the action depends on the type
	switch a.Type {
	// if we are retrieving the information for a given repo
	case common.ActionBuildProject:

		fmt.Println(a.Payload)
	}
}
