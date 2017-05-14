package objectTypes

import (
	"github.com/graphql-go/graphql"

	"github.com/graphql-go/relay"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "MaestroAPI",
	Fields: graphql.Fields{
		"logs": &graphql.Field{
			Type:    LogEntryDefinition.ConnectionType,
			Args:    relay.ConnectionArgs,
			Resolve: ResolveLogEntrys,
		},
		"project": &graphql.Field{
			Type:    Project,
			Resolve: ResolveProject,
		},
	},
})
