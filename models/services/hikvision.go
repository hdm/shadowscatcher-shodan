package services

type Hikvision struct {
	ActivexFiles  map[string]string `json:"activex_files"`
	WebVersion    string            `json:"web_version"`
	CustomVersion string            `json:"custom_version"`
	DeviceVersion string            `json:"device_version"`
	DeviceModel   string            `json:"device_model"`
	DeviceName    string            `json:"device_name"`
	PluginVersion string            `json:"plugin_version"`
}
