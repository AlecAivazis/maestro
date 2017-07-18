package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecaivazis/maestro/common"
	"github.com/nautilus/events"
)

// MaestroProjects is the service responsible for projects and their tickets, as well as
// kicking off builds at the correct time.
type MaestroProjects struct {
	events.EventBroker
}

// for now, just keep the projects in memory
var projects = map[string]Project{
	"hello": &ProjectInMemory{"Hello", []Ticket{
		&TicketInMemory{
			name: "Hello",
			status: &StatusInMemory{
				name: "Ready",
			},
		},
	}},
}

func (s *MaestroProjects) HandleAction(a *events.Action) {
	// respond to the appropriate action
	switch a.Type {
	// if we are looking to retrieve a project from the storage
	case common.ActionRetrieveProject:
		// unmarshal the payload into something we understand
		payload := common.RetrieveProjectPayload{}
		err := json.Unmarshal([]byte(a.Payload), &payload)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// if we know of the project being requested
		if proj, ok := projects[payload.Name]; ok {
			// marshal the project we found
			marshaled, err := json.Marshal(common.Project{
				Name: proj.Name(),
			})
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// the action returning the requested information
			events.Reply(s, a, &events.Action{
				Type:    "Reply",
				Payload: string(marshaled),
			})
			// otherwise the payload is asking for a project we do not recognize
		} else {
			// we have to reply with an error but I dont know how to do that yet in go
			fmt.Println("could not find project with name", payload.Name)
		}
	}
}
