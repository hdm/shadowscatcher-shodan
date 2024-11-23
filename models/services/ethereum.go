package services

type EthereumRPC struct {
	ChainID  string   `json:"chain_id,omitempty"`
	Client   string   `json:"client,omitempty"`
	Version  string   `json:"version,omitempty"`
	Accounts []string `json:"accounts,omitempty"`
	Hashrate string   `json:"hashrate,omitempty"`
	Compiler string   `json:"compiler,omitempty"`
	Platform string   `json:"platform,omitempty"`
}

type EthereumP2P struct {
	Neighbors []struct {
		IP      string `json:"ip,omitempty"`
		Pubkey  string `json:"pubkey,omitempty"`
		UDPPort int    `json:"udp_port,omitempty"`
		TCPPort int    `json:"tcp_port,omitempty"`
	} `json:"neighbors,omitempty"`
	TCPPort int    `json:"tcp_port,omitempty"`
	Pubkey  string `json:"pubkey,omitempty"`
	UDPPort int    `json:"udp_port,omitempty"`
}
