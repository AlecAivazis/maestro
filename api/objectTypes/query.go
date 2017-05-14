package objectTypes

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	gql "github.com/nautilus/api"
	"github.com/nautilus/events"

	"github.com/AlecAivazis/maestro/common"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "MaestroAPI",
	Fields: graphql.Fields{
		"logs": &graphql.Field{
			Type: LogEntryDefinition.ConnectionType,
			Args: relay.ConnectionArgs,
			Resolve: gql.ConnectionResolver(&gql.RemoteObjectSpec{
				Service: "log",
				Action: func(p graphql.ResolveParams) (*events.Action, error) {
					// marshal the request
					req, err := json.Marshal(common.RetrieveLogPayload{
						Label: "BuildProject",
					})
					if err != nil {
						fmt.Println(err.Error())
						return nil, err
					}

					// publish an action
					return &events.Action{
						Type:    common.ActionRetrieveLogs,
						Payload: string(req),
					}, nil
				},
			}),
		},
		"project": &graphql.Field{
			Type: Project,
			Resolve: gql.ObjectResolver(&gql.RemoteObjectSpec{
				Service: "projects",
				Action: func(p graphql.ResolveParams) (*events.Action, error) {
					// marshal the request
					req, err := json.Marshal(common.RetrieveProjectPayload{
						Name: "hello",
					})
					if err != nil {
						fmt.Println(err.Error())
						return nil, err
					}

					return &events.Action{
						Type:    common.ActionRetrieveProject,
						Payload: string(req),
					}, nil
				},
			}),
		},
	},
})
