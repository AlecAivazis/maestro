package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
	"github.com/nautilus/events"
	"github.com/spf13/afero"

	"github.com/AlecAivazis/maestro/common"
)

var (
	cli                  *client.Client
	golangDockerTemplate *template.Template
)

// MaestroRepo is the service responsible for retrieving information for a
// given repo
type MaestroRepo struct {
	events.EventBroker
	Fs afero.Fs
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
		doc := &bytes.Buffer{}
		err = golangDockerTemplate.Execute(doc, payload)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// we should publish the output to the log service
		writer, err := common.LogWriter(s, "BuildProject")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// create a tarball
		tempDir := afero.GetTempDir(s.Fs, "")
		tarballPath := filepath.Join(tempDir, "docker")
		tar := new(archivex.TarFile)
		tar.Create(tarballPath)
		tar.Add("Dockerfile", doc.Bytes())
		tar.Close()

		dockerBuildContext, err := os.Open(tarballPath + ".tar")
		defer dockerBuildContext.Close()

		opts := types.ImageBuildOptions{}

		resp, err := cli.ImageBuild(context.Background(), dockerBuildContext, opts)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// make sure we close the body when we're done
		defer resp.Body.Close()

		// copy the output to the log service
		if _, err = io.Copy(writer, resp.Body); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func init() {
	// connect to the local docker socket to build the image
	dClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	cli = dClient

	// compile the dockerfile templates
	golangDockerTemplate = template.Must(template.New("golang").Parse(dockerfileGolang))
}
