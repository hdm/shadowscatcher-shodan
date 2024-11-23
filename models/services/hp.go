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

type HPPrinterWeb struct {
	UUID                          string `json:"uuid,omitempty"`
	MakeAndModelFamily            string `json:"make_and_model_family,omitempty"`
	SkuIdentifier                 string `json:"sku_identifier,omitempty"`
	PasswordStatus                string `json:"password_status,omitempty"`
	ProductDerivativeNumber       string `json:"product_derivative_number,omitempty"`
	ProductDerivativeFlavorStatus string `json:"product_derivative_flavor_status,omitempty"`
	MakeAndModel                  string `json:"make_and_model,omitempty"`
	NetworkSummary                struct {
		EthernetAdapter struct {
			Name                       string `json:"name,omitempty"`
			Power                      string `json:"power,omitempty"`
			SpeedDuplexNegotiationMode string `json:"speed_duplex_negotiation_mode,omitempty"`
			IsCurrentRequestInterface  string `json:"is_current_request_interface,omitempty"`
			DeviceConnectivityPortType string `json:"device_connectivity_port_type,omitempty"`
			IsConnected                string `json:"is_connected,omitempty"`
			AdminDisabled              string `json:"admin_disabled,omitempty"`
		} `json:"ethernet_adapter,omitempty"`
		WifiAdapter struct {
			Name                       string `json:"name,omitempty"`
			Power                      string `json:"power,omitempty"`
			IsCurrentRequestInterface  string `json:"is_current_request_interface,omitempty"`
			DeviceConnectivityPortType string `json:"device_connectivity_port_type,omitempty"`
			IsConnected                string `json:"is_connected,omitempty"`
			AdminDisabled              string `json:"admin_disabled,omitempty"`
		} `json:"wifi_adapter,omitempty"`
		UsbAdapter struct {
			DeviceConnectivityPortType string `json:"device_connectivity_port_type,omitempty"`
			IsConnected                string `json:"is_connected,omitempty"`
			Name                       string `json:"name,omitempty"`
			Power                      string `json:"power,omitempty"`
		} `json:"usb_adapter,omitempty"`
		AccessPointAdapter struct {
			Name                       string `json:"name,omitempty"`
			Power                      string `json:"power,omitempty"`
			IsCurrentRequestInterface  string `json:"is_current_request_interface,omitempty"`
			DeviceConnectivityPortType string `json:"device_connectivity_port_type,omitempty"`
			IsConnected                string `json:"is_connected,omitempty"`
			AdminDisabled              string `json:"admin_disabled,omitempty"`
		} `json:"access_point_adapter,omitempty"`
	} `json:"network_summary,omitempty"`
	DuplexUnit         string `json:"duplex_unit,omitempty"`
	ProductNumber      string `json:"product_number,omitempty"`
	WoobeseriesID      string `json:"woobeseries_id,omitempty"`
	SerialNumber       string `json:"serial_number,omitempty"`
	ServiceID          string `json:"service_id,omitempty"`
	FirmwareDownloadID string `json:"firmware_download_id,omitempty"`
	VersionDate        string `json:"version_date,omitempty"`
	Revision           string `json:"revision,omitempty"`
	MakeAndModelBase   string `json:"make_and_model_base,omitempty"`
}
