package services

type AnyConnect struct {
	ConfigHash  string `json:"config_hash,omitempty"`
	TunnelGroup string `json:"tunnel_group,omitempty"`
	GroupAlias  string `json:"group_alias,omitempty"`
}
