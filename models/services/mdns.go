package services

type MDNS struct {
	Services map[string]struct {
		Port int      `json:"port,omitempty"`
		Data []string `json:"data,omitempty"`
		Ptr  string   `json:"ptr,omitempty"`
		Name string   `json:"name,omitempty"`
		IPv4 []string `json:"ipv4,omitempty"`
		IPv6 []string `json:"ipv6,omitempty"`
	} `json:"services,omitempty"`
	Answers struct {
		Ptr []string `json:"PTR,omitempty"`
	} `json:"answers,omitempty"`
}
