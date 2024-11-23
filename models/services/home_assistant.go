package services

type HomeAssistant struct {
	InternalURL      string `json:"internal_url,omitempty"`
	LocationName     string `json:"location_name,omitempty"`
	UUID             string `json:"uuid,omitempty"`
	InstallationType string `json:"installation_type,omitempty"`
	BaseURL          string `json:"base_url,omitempty"`
}
