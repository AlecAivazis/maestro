package common

type BuildProjectPayload struct {
	URL    string
	Branch string
}

const (
	ActionBuildProject = "BuildProject"
)
