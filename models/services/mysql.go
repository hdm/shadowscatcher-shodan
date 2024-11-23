package services

type MySQL struct {
	AuthenticationPlugin       string `json:"authentication_plugin"`
	Capabilities               int    `json:"capabilities"`
	ServerStatus               int    `json:"server_status"`
	Version                    string `json:"version"`
	ErrorMessage               string `json:"error_message"`
	ErrorCode                  int    `json:"error_code"`
	ExtendedServerCapabilities int    `json:"extended_server_capabilities"`
	ProtocolVersion            int    `json:"protocol_version"`
	ServerLanguage             int    `json:"server_language"`
}
