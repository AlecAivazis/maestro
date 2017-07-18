package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nautilus/events"

	"github.com/AlecAivazis/maestro/common"
)

// MaestroLogging is the service responsible for collecting and aggregating logs.
type MaestroLogging struct {
	events.EventBroker
}

var logCache = make(map[string][]common.LogEntry)

func (s *MaestroLogging) HandleAction(a *events.Action) {
	// what we do with the action depends on the type
	switch a.Type {
	// if we are logging something new
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
		// the log entry for the input
		entry := common.LogEntry{
			Body:        payload.Payload,
			DateCreated: time.Now().Format("Mon Jan _2 15:04:05 2006"),
		}

		// look for logs with the current label
		val, ok := logCache[payload.Label]
		// if there is an entry in the cache
		if ok {
			// add the payload to the cache
			logCache[payload.Label] = append(val, entry)
		} else {
			// otherwise there is no log with that label so save one
			logCache[payload.Label] = []common.LogEntry{entry}
		}

	// if we need to retrieve logs for a particular project
	case common.ActionRetrieveLogs:
		// the request for logs
		req := common.RetrieveLogPayload{}

		// try to unmarhsal the event payload into the structure we want
		err := json.Unmarshal([]byte(a.Payload), &req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// marshal the appropriate log entries
		str, err := json.Marshal(logCache[req.Label])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// send the logs to the requester
		events.Reply(s, a, &events.Action{
			Type:    "Hello",
			Payload: string(str),
		})
	}
}
