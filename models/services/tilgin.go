package services

type TilginHomeGateway struct {
	SoftwareFamily   string `json:"software_family,omitempty"`
	SoftwareRevision string `json:"software_revision,omitempty"`
	ProductName      string `json:"product_name,omitempty"`
}
