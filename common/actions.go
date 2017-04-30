package common

type BuildProjectPayload struct {
	URL     string `json:"url"`
	Branch  string `json:"branch"`
	TempDir string `json:"tempDir"`
}

const (
	ActionBuildProject = "BuildProject"
	ActionLogAction    = "LogAction"
)
