package services

type ADB struct {
	Device string `json:"device,omitempty"`
	Model  string `json:"model,omitempty"`
	Name   string `json:"name,omitempty"`
}
