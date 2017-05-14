package common

type BuildProjectPayload struct {
	URL     string `json:"url"`
	Branch  string `json:"branch"`
	TempDir string `json:"tempDir"`
}

type LogPayload struct {
	Label   string `json:"label"`
	Payload string `json:"payload"`
}

type RetrieveProjectPayload struct {
	Name string `json:"name"`
}

type RetrieveLogPayload struct {
	Label string `json:"label"`
}

const (
	ActionBuildProject    = "BuildProject"
	ActionLogAction       = "LogAction"
	ActionRetrieveLogs    = "RetrieveLogs"
	ActionRetrieveProject = "RetrieveProject"
)
