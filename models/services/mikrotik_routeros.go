package services

type MikrotikRouterOS struct {
	Interfaces []string `json:"interfaces,omitempty"`
	Version    string   `json:"version"`
}
