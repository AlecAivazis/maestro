package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/fsouza/go-dockerclient"
	"github.com/nautilus/events"

	"github.com/AlecAivazis/maestro/common"
)

var (
	client               *docker.Client
	golangDockerTemplate *template.Template
)

// MaestroRepo is the service responsible for retrieving information for a
// given repo
type MaestroRepo struct {
	events.EventBroker
}

func (s *MaestroRepo) HandleAction(a *events.Action) {
	// what we do with the action depends on the type
	switch a.Type {
	// if we are retrieving the information for a given repo
	case common.ActionBuildProject:
		// unmarshal the payload into something we understand
		payload := common.BuildProjectPayload{}
		err := json.Unmarshal([]byte(a.Payload), &payload)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// render the template with the appropriate language
		dockerfile := bytes.Buffer{}
		err = golangDockerTemplate.Execute(&dockerfile, payload)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// build the image in the directory
		// opts := docker.BuildImageOptions{
		// 	Name:        payload.Branch,
		// 	InputStream: &dockerfile,
		// }
		// run the maestro script inside of the container

	}
}

func init() {
	// connect to the local docker socket to build the image
	endpoint := "unix:///var/run/docker.sock"
	dClient, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	client = dClient

	// compile the dockerfile templates
	golangDockerTemplate = template.Must(template.New("golang").Parse(dockerfileGolang))
}
