package objectTypes

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var Ticket = graphql.NewObject(graphql.ObjectConfig{
	Name: "Ticket",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var TicketDefinition = relay.ConnectionDefinitions(relay.ConnectionConfig{
	Name:     "Ticket",
	NodeType: Ticket,
})
