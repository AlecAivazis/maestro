package objectTypes

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
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
