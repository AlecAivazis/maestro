package objectTypes

import (
	"github.com/graphql-go/graphql"
	"github.com/nautilus/events"
	"github.com/nautilus/services/graphql"

	"encoding/json"

	"github.com/AlecAivazis/maestro/common"
	"github.com/graphql-go/relay"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "MaestroAPI",
	Fields: graphql.Fields{
		"logs": &graphql.Field{
			Type: LogEntryDefinition.ConnectionType,
			Args: relay.ConnectionArgs,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// the list of logs
				logs := []interface{}{}

				// the broker that made the request
				broker := p.Context.Value(GraphqlService.BrokerCtx).(events.EventBroker)
				// a channel to recieve a response
				ansChan := make(chan *events.Action, 1)
				errChan := make(chan error, 1)

				// publish an action
				broker.Ask("log", ansChan, errChan, &events.Action{
					Type:    common.ActionRetrieveLog,
					Payload: "BuildProject",
				})

				// wait for some kind of a reply
				select {
				// if we were successful
				case r := <-ansChan:
					// treat the payload like json
					err := json.Unmarshal([]byte(r.Payload), &logs)
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

				// convert args map[string]interface into ConnectionArguments
				args := relay.NewConnectionArguments(p.Args)

				// return a connection from the logs array
				return relay.ConnectionFromArray(logs, args), nil
			},
		},
	},
})
