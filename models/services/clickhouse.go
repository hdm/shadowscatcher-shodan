package services

type ClickHouse struct {
	PrimaryStatus  string `json:"primary_status,omitempty"`
	RequiredLogin  bool   `json:"required_login,omitempty"`
	Version        string `json:"version,omitempty"`
	ServerName     string `json:"server_name,omitempty"`
	ReplicasStatus string `json:"replicas_status,omitempty"`
}
