package mutations

import (
	"golang.org/x/net/context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var triggerBuild = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "TriggerBuild",
	InputFields: graphql.InputObjectConfigFieldMap{
		"repo": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The name of the repo to build from",
		},
		"branch": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The branch of the repo to build from",
		},
	},
	OutputFields: graphql.Fields{
		"success": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
	MutateAndGetPayload: func(input map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
		return map[string]interface{}{
			"success": true,
		}, nil
	},
})
