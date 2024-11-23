package services

type SynologyDSM struct {
	Hostname         string `json:"hostname,omitempty"`
	CustomLoginTitle string `json:"custom_login_title,omitempty"`
}

type SynologySRM struct {
	Hostname string `json:"hostname,omitempty"`
}
