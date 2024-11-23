package models

type MAC struct {
	Date       string `json:"date,omitempty"`
	Org        string `json:"org,omitempty"`
	Oui        string `json:"oui,omitempty"`
	Assignment string `json:"assignment,omitempty"`
}
