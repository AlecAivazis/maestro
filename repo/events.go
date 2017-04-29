package main

import (
	"fmt"

	"github.com/nautilus/events"

	"github.com/AlecAivazis/maestro/common"
)

func (s *MaestroRepo) HandleAction(a *events.Action) {
	// what we do with the action depends on the type
	switch a.Type {
	// if we are retrieving the information for a given repo
	case common.ActionRetrieveRepo:
		// get the information associated with the particular repo
		reply := &events.Action{
			Type:    "Awesome",
			Payload: "Opossum",
		}

		// send the reply
		err := events.Reply(s, a, reply)
		if err != nil {
			fmt.Println(err)
		}
	}
}
