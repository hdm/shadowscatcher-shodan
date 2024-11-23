package services

type PPTP struct {
	Firmware int    `json:"firmware,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Hostname string `json:"hostname,omitempty"`
}
