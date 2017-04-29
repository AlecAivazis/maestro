package objectTypes

import (
	"github.com/graphql-go/graphql"
	"github.com/nautilus/events"
	"github.com/nautilus/services/graphql"

	"github.com/AlecAivazis/maestro/common"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "MaestroAPI",
	Fields: graphql.Fields{
		"viewer": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// the broker that made the request
				broker := p.Context.Value(GraphqlService.BrokerCtx).(events.EventBroker)
				// a channel to recieve a response
				ansChan := make(chan string, 1)
				errChan := make(chan error, 1)

				// publish an action
				broker.Ask("repo", &events.Action{
					Type:    common.ActionRetrieveRepo,
					Payload: "world",
				}, ansChan, errChan)

				// wait for some kind of a reply
				select {
				// if we were successful
				case r := <-ansChan:
					// return the response
					return r, nil
				// if something went wrong
				case e := <-errChan:
					// return the error
					return nil, e
				}
			},
		},
	},
})
