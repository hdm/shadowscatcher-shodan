package services

type Fortinet struct {
	Device       string `json:"device,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
	Model        string `json:"model,omitempty"`
}
