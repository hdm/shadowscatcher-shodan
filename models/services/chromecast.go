package services

type Chromecast struct {
	DeviceInfo struct {
		PublicKey     string `json:"public_key,omitempty"`
		UmaClientID   string `json:"uma_client_id,omitempty"`
		ProductName   string `json:"product_name,omitempty"`
		HotspotBssid  string `json:"hotspot_bssid,omitempty"`
		DeviceName    string `json:"device_name,omitempty"`
		MacAddress    string `json:"mac_address,omitempty"`
		SsdpUdn       string `json:"ssdp_udn,omitempty"`
		CloudDeviceID string `json:"cloud_device_id,omitempty"`
		ModelName     string `json:"model_name,omitempty"`
		Manufacturer  string `json:"manufacturer,omitempty"`
	} `json:"device_info,omitempty"`
	Wifi struct {
		Ssid  string `json:"ssid,omitempty"`
		Bssid string `json:"bssid,omitempty"`
	} `json:"wifi,omitempty"`
	Net struct {
		EthernetConnected bool   `json:"ethernet_connected,omitempty"`
		IPAddress         string `json:"ip_address,omitempty"`
		Online            bool   `json:"online,omitempty"`
	} `json:"net,omitempty"`
	Version   int `json:"version,omitempty"`
	BuildInfo struct {
		SystemBuildNumber  string `json:"system_build_number,omitempty"`
		CastControlVersion int    `json:"cast_control_version,omitempty"`
		CastBuildRevision  string `json:"cast_build_revision,omitempty"`
		ReleaseTrack       string `json:"release_track,omitempty"`
		BuildType          int    `json:"build_type,omitempty"`
	} `json:"build_info,omitempty"`
}
