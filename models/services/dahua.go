package services

type Dahua struct {
	SerialNumber string `json:"serial_number,omitempty"`
}

type DahuaDVRWeb struct {
	WebVersion string `json:"web_version,omitempty"`
	Plugin     struct {
		Classid    string `json:"classid,omitempty"`
		Version    string `json:"version,omitempty"`
		Name       string `json:"name,omitempty"`
		MacVersion string `json:"mac_version,omitempty"`
	} `json:"plugin,omitempty"`
	ChannelNames []string `json:"channel_names,omitempty"`
}
