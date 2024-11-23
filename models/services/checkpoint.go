package services

type Checkpoint struct {
	FirewallHost    string `json:"firewall_host,omitempty"`
	SmartcenterHost string `json:"smartcenter_host,omitempty"`
}
