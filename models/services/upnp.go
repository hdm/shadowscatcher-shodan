package services

type UPNP struct {
	ModelURL         string `json:"model_url,omitempty"`
	ManufacturerURL  string `json:"manufacturer_url,omitempty"`
	FriendlyName     string `json:"friendly_name,omitempty"`
	ModelNumber      string `json:"model_number,omitempty"`
	PresentationURL  string `json:"presentation_url,omitempty"`
	DeviceType       string `json:"device_type,omitempty"`
	ModelDescription string `json:"model_description,omitempty"`
	SerialNumber     string `json:"serial_number,omitempty"`
	Udn              string `json:"udn,omitempty"`
	ModelName        string `json:"model_name,omitempty"`
	Manufacturer     string `json:"manufacturer,omitempty"`
}
