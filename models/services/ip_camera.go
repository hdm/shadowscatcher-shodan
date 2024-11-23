package services

type IPCamera struct {
	Name          string `json:"name,omitempty"`
	AliasName     string `json:"alias_name,omitempty"`
	AppVersion    string `json:"app_version,omitempty"`
	SystemVersion string `json:"system_version,omitempty"`
	ID            string `json:"id,omitempty"`
	DdnsHost      string `json:"ddns_host,omitempty"`
}
