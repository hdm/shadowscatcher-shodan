package services

import "encoding/json"

type Bitcoin struct {
	Handshake []struct {
		Nonce    json.Number `json:"nonce,omitempty"`
		Relay    bool        `json:"relay,omitempty"`
		FromAddr struct {
			Services  int    `json:"services,omitempty"`
			Timestamp any    `json:"timestamp,omitempty"`
			IPv4      string `json:"ipv4,omitempty"`
			Port      int    `json:"port,omitempty"`
			IPv6      string `json:"ipv6,omitempty"`
		} `json:"from_addr,omitempty"`
		Checksum    string `json:"checksum,omitempty"`
		MagicNumber string `json:"magic_number,omitempty"`
		ToAddr      struct {
			Services  int    `json:"services,omitempty"`
			Timestamp any    `json:"timestamp,omitempty"`
			IPv4      string `json:"ipv4,omitempty"`
			Port      int    `json:"port,omitempty"`
			IPv6      string `json:"ipv6,omitempty"`
		} `json:"to_addr,omitempty"`
		Lastblock int    `json:"lastblock,omitempty"`
		Length    int    `json:"length,omitempty"`
		Version   int    `json:"version,omitempty"`
		Command   string `json:"command,omitempty"`
		UserAgent string `json:"user_agent,omitempty"`
		Timestamp int    `json:"timestamp,omitempty"`
		Services  int    `json:"services,omitempty"`
	} `json:"handshake,omitempty"`
	Addresses []any `json:"addresses,omitempty"`
}
