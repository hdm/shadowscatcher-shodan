package services

type VNC struct {
	ProtocolVersion string            `json:"protocol_version,omitempty"`
	SecurityTypes   map[string]string `json:"security_types,omitempty"`
}
