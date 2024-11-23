package services

type QNAP struct {
	Model struct {
		DisplayModelName  string `json:"display_model_name,omitempty"`
		Platform          string `json:"platform,omitempty"`
		InternalModelName string `json:"internal_model_name,omitempty"`
		PlatformEx        string `json:"platform_ex,omitempty"`
		ModelName         string `json:"model_name,omitempty"`
	} `json:"model,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Firmware struct {
		Version string `json:"version,omitempty"`
		Number  string `json:"number,omitempty"`
		Build   string `json:"build,omitempty"`
	} `json:"firmware,omitempty"`
	Apps map[string]map[string]string `json:"apps,omitempty"`
}
