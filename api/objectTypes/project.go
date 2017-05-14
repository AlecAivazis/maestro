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

var ResolveProject = gql.ObjectResolver(&gql.RemoteObjectSpec{
	Service: "projects",
	WithAction: func(p graphql.ResolveParams) (*events.Action, error) {
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
})
