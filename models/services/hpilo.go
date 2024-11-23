package services

type HPiLO struct {
	IloType    string `json:"ilo_type,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	ServerType string `json:"server_type,omitempty"`
	IloUUID    string `json:"ilo_uuid,omitempty"`
	Nics       []struct {
		Status      string `json:"status,omitempty"`
		Description string `json:"description,omitempty"`
		Location    string `json:"location,omitempty"`
		MacAddress  string `json:"mac_address,omitempty"`
		IPAddress   string `json:"ip_address,omitempty"`
		Port        string `json:"port,omitempty"`
	} `json:"nics,omitempty"`
	IloFirmware     string `json:"ilo_firmware,omitempty"`
	IloSerialNumber string `json:"ilo_serial_number,omitempty"`
	SerialNumber    string `json:"serial_number,omitempty"`
	Cuuid           string `json:"cuuid,omitempty"`
	ProductID       string `json:"product_id,omitempty"`
}
