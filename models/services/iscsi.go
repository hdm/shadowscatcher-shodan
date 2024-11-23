package services

type ISCSI struct {
	Targets []struct {
		Name      string   `json:"name,omitempty"`
		Addresses []string `json:"addresses,omitempty"`
	} `json:"targets,omitempty"`
}
