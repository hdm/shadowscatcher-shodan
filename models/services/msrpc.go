package services

type MSRPC struct {
	Towers map[string]struct {
		Bindings []struct {
			NcacnIPTCP string `json:"ncacn_ip_tcp,omitempty"`
			Ncalrpc    string `json:"ncalrpc,omitempty"`
			NcacnNp    string `json:"ncacn_np,omitempty"`
			Netbios    string `json:"netbios,omitempty"`
		} `json:"bindings,omitempty"`
		Version    string `json:"version,omitempty"`
		Annotation string `json:"annotation,omitempty"`
		Provider   string `json:"provider,omitempty"`
		Protocol   string `json:"protocol,omitempty"`
	} `json:"towers,omitempty"`
	ActualCount int `json:"actual_count,omitempty"`
}
