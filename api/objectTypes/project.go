package objectTypes

import (
	"errors"

	"github.com/AlecAivazis/maestro/common"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var Project = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"tickets": &graphql.Field{
			Type: TicketDefinition.ConnectionType,
			Args: relay.ConnectionArgs,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// try to turn the source into a project
				if proj, ok := p.Source.(common.Project); ok {
					// convert args map[string]interface into ConnectionArguments
					args := relay.NewConnectionArguments(p.Args)

					// copy the tickets to a slice of interface
					tickets := make([]interface{}, len(proj.Tickets))
					for i, v := range proj.Tickets {
						tickets[i] = v
					}

					// resolve the list of tickets for the project
					return relay.ConnectionFromArray(tickets, args), nil
				}

				return nil, errors.New("project was not resolved with a project object")
			},
		},
	},
})

var ProjectDefinition = relay.ConnectionDefinitions(relay.ConnectionConfig{
	Name:     "Project",
	NodeType: Project,
})
