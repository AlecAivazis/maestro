package objectTypes

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var LogEntry = graphql.NewObject(graphql.ObjectConfig{
	Name: "LogEntry",
	Fields: graphql.Fields{
		"dateCreated": &graphql.Field{
			Type: graphql.String,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var LogEntryDefinition = relay.ConnectionDefinitions(relay.ConnectionConfig{
	Name:     "Log",
	NodeType: LogEntry,
})
