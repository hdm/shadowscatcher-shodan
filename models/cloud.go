package models

type Cloud struct {
	Region   string `json:"region"`
	Service  string `json:"service"`
	Provider string `json:"provider"`
}
