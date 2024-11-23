package services

type Hikvision struct {
	ActivexFiles  map[string]string `json:"activex_files"`
	WebVersion    string            `json:"web_version"`
	CustomVersion string            `json:"custom_version"`
	DeviceVersion string            `json:"device_version"`
	PluginVersion string            `json:"plugin_version"`
}
