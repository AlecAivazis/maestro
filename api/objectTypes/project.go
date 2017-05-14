package objectTypes

import (
	"encoding/json"
	"fmt"

	"github.com/AlecAivazis/maestro/common"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/nautilus/events"
	"github.com/nautilus/services/graphql"
)

var Project = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var ProjectDefinition = relay.ConnectionDefinitions(relay.ConnectionConfig{
	Name:     "Project",
	NodeType: Project,
})

var ResolveProject = func(p graphql.ResolveParams) (interface{}, error) {
	// the list of logs
	var project interface{}

	// the broker that made the request
	broker := p.Context.Value(GraphqlService.BrokerCtx).(events.EventBroker)
	// a channel to recieve a response
	ansChan := make(chan *events.Action, 1)
	errChan := make(chan error, 1)

	// marshal the request
	req, err := json.Marshal(common.RetrieveProjectPayload{
		Name: "hello",
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// publish an action
	broker.Ask("projects", ansChan, errChan, &events.Action{
		Type:    common.ActionRetrieveProject,
		Payload: string(req),
	})

	// wait for some kind of a reply
	select {
	// if we were successful
	case r := <-ansChan:
		// treat the payload like json
		err := json.Unmarshal([]byte(r.Payload), &project)
		// if somthing went wrong
		if err != nil {
			// return the error
			return nil, err
		}
	// if something went wrong
	case e := <-errChan:
		// return the error
		return nil, e
	}

	// return a connection from the logs array
	return project, nil
}
