package services

type Airplay struct {
	FirmwareBuild     string `json:"firmware_build,omitempty"`
	Name              string `json:"name,omitempty"`
	DeviceModel       string `json:"device_model,omitempty"`
	HardwareRevision  string `json:"hardware_revision,omitempty"`
	ProtocolVersion   string `json:"protocol_version,omitempty"`
	SerialNumber      string `json:"serial_number,omitempty"`
	FirmwareBuildDate string `json:"firmware_build_date,omitempty"`
	Sdk               string `json:"sdk,omitempty"`
	DeviceID          string `json:"device_id,omitempty"`
}
