package services

type MySQLX struct {
	TLS               bool `json:"tls,omitempty"`
	ClientPwdExpireOk bool `json:"client.pwd_expire_ok,omitempty"`
	Compression       struct {
		Algorithm []string `json:"algorithm,omitempty"`
	} `json:"compression,omitempty"`
	DocFormats               string   `json:"doc.formats,omitempty"`
	NodeType                 string   `json:"node_type,omitempty"`
	ClientInteractive        bool     `json:"client.interactive,omitempty"`
	AuthenticationMechanisms []string `json:"authentication.mechanisms,omitempty"`
}
