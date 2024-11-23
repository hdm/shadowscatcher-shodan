package services

type NTLM struct {
	TargetRealm         string   `json:"target_realm,omitempty"`
	DNSDomainName       string   `json:"dns_domain_name,omitempty"`
	DNSForestName       string   `json:"dns_forest_name,omitempty"`
	NetbiosDomainName   string   `json:"netbios_domain_name,omitempty"`
	Timestamp           int64    `json:"timestamp,omitempty"`
	FQDN                string   `json:"fqdn,omitempty"`
	NetbiosComputerName string   `json:"netbios_computer_name,omitempty"`
	OS                  []string `json:"os,omitempty"`
	OSBuild             string   `json:"os_build,omitempty"`
}
