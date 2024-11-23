package services

type BACnet struct {
	Object     string `json:"object,omitempty"`
	Vendor     string `json:"vendor,omitempty"`
	Name       string `json:"name,omitempty"`
	Firmware   string `json:"firmware,omitempty"`
	InstanceID string `json:"instance_id,omitempty"`
	Bbmd       []any  `json:"bbmd,omitempty"`
	Fdt        []any  `json:"fdt,omitempty"`
	Model      string `json:"model,omitempty"`
	Desc       string `json:"desc,omitempty"`
	AppSoft    string `json:"appsoft,omitempty"`
}
