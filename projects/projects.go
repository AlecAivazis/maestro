package main

import (
	"github.com/nautilus/events"
)

// MaestroProjects is the service responsible for projects and their tickets, as well as
// kicking off builds at the correct time.
type MaestroProjects struct {
	events.EventBroker
}

func (s *MaestroProjects) HandleAction(a *events.Action) {
}
