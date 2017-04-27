package objectTypes

import "github.com/graphql-go/graphql"

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "MaestroAPI",
	Fields: graphql.Fields{
		"viewer": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "hello", nil
			},
		},
	},
})
