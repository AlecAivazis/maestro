package main

import (
	"fmt"

	"github.com/nautilus/events"

	"encoding/json"

	"github.com/AlecAivazis/maestro/common"
)

// MaestroLogging is the service responsible for collecting and aggregating logs.
type MaestroLogging struct {
	events.EventBroker
}

var logCache = make(map[string][]string)

func (s *MaestroLogging) HandleAction(a *events.Action) {
	// what we do with the action depends on the type
	switch a.Type {
	// if we are retrieving the information for a given repo
	case common.ActionLogAction:
		// a log payload to write to
		payload := common.LogPayload{}

		// try to parse the payload into the struct we know it to follow
		err := json.Unmarshal([]byte(a.Payload), &payload)
		// if something went wrong
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// look for logs with the current label
		val, ok := logCache[payload.Label]
		// if there is an entry in the cache
		if ok {
			// add the payload to the cache
			logCache[payload.Label] = append(val, payload.Payload)
		} else {
			// otherwise there is no log with that label so save one
			logCache[payload.Label] = []string{payload.Payload}
		}
	}
}
