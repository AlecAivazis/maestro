package mutations

import (
	"golang.org/x/net/context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/nautilus/events"
	"github.com/nautilus/services/graphql"

	"encoding/json"

	"github.com/AlecAivazis/maestro/common"
)

var triggerBuild = relay.MutationWithClientMutationID(relay.MutationConfig{
	Name: "TriggerBuild",
	InputFields: graphql.InputObjectConfigFieldMap{
		"url": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The url of the repo to build from",
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
		// the broker that made the request
		broker := ctx.Value(GraphqlService.BrokerCtx).(events.EventBroker)

		// the input payload
		buildInput, err := json.Marshal(&common.BuildProjectPayload{
			URL:    input["url"].(string),
			Branch: input["branch"].(string),
		})
		if err != nil {
			return nil, err
		}

		// publish the build project event
		err = broker.Publish("build", &events.Action{
			Type:    common.ActionBuildProject,
			Payload: string(buildInput),
		})

		// if there was something wrong
		if err != nil {
			return map[string]interface{}{
				"success": false,
			}, err
		}
		// nothing went wrong
		return map[string]interface{}{
			"success": true,
		}, nil
	},
})
