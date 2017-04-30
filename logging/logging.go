package main

import (
	"fmt"

	"github.com/nautilus/events"

	"github.com/AlecAivazis/maestro/common"
)

// MaestroLogging is the service responsible for collecting and aggregating logs.
type MaestroLogging struct {
	events.EventBroker
}

func (s *MaestroLogging) HandleAction(a *events.Action) {
	// what we do with the action depends on the type
	switch a.Type {
	// if we are retrieving the information for a given repo
	case common.ActionLogAction:
		fmt.Println(a.Payload)
	}
}
