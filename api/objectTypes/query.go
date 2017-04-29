package objectTypes

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/nautilus/events"
	"github.com/nautilus/services/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "MaestroAPI",
	Fields: graphql.Fields{
		"viewer": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// the broker that made the request
				broker := p.Context.Value(GraphqlService.BrokerCtx)
				// assert that its an event broker
				if broker, ok := broker.(events.EventBroker); ok {
					fmt.Println(broker)
				}
				return "hello", nil
			},
		},
	},
})
