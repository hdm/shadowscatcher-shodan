package services

import "net"

type Stun struct {
	ServerIP net.IP `json:"server_ip"`
	Software string `json:"software"`
}
