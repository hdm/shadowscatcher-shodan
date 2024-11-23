package services

type NodeExporter struct {
	NodeExporterBuildInfo struct {
		Tags      string `json:"tags,omitempty"`
		Version   string `json:"version,omitempty"`
		Goarch    string `json:"goarch,omitempty"`
		Branch    string `json:"branch,omitempty"`
		Goos      string `json:"goos,omitempty"`
		Revision  string `json:"revision,omitempty"`
		Goversion string `json:"goversion,omitempty"`
	} `json:"node_exporter_build_info,omitempty"`
	NodeUnameInfo struct {
		Sysname    string `json:"sysname,omitempty"`
		Nodename   string `json:"nodename,omitempty"`
		Domainname string `json:"domainname,omitempty"`
		Machine    string `json:"machine,omitempty"`
		Version    string `json:"version,omitempty"`
		Release    string `json:"release,omitempty"`
	} `json:"node_uname_info,omitempty"`
	NodeOsInfo struct {
		Name            string `json:"name,omitempty"`
		PrettyName      string `json:"pretty_name,omitempty"`
		IDLike          string `json:"id_like,omitempty"`
		VersionID       string `json:"version_id,omitempty"`
		VersionCodename string `json:"version_codename,omitempty"`
		Version         string `json:"version,omitempty"`
		ID              string `json:"id,omitempty"`
	} `json:"node_os_info,omitempty"`
	NodeDmiInfo struct {
		BiosVendor      string `json:"bios_vendor,omitempty"`
		ProductVersion  string `json:"product_version,omitempty"`
		SystemVendor    string `json:"system_vendor,omitempty"`
		ProductName     string `json:"product_name,omitempty"`
		BiosDate        string `json:"bios_date,omitempty"`
		BiosVersion     string `json:"bios_version,omitempty"`
		BiosRelease     string `json:"bios_release,omitempty"`
		ChassisVendor   string `json:"chassis_vendor,omitempty"`
		ChassisAssetTag string `json:"chassis_asset_tag,omitempty"`
		ChassisVersion  string `json:"chassis_version,omitempty"`
		ChassisSerial   string `json:"chassis_serial,omitempty"`
		ProductUUID     string `json:"product_uuid,omitempty"`
		ProductSerial   string `json:"product_serial,omitempty"`
		ProductFamily   string `json:"product_family,omitempty"`
		ProductSKU      string `json:"product_sku,omitempty"`
		BoardAssetTag   string `json:"board_asset_tag,omitempty"`
		BoardName       string `json:"board_name,omitempty"`
		BoardVendor     string `json:"board_vendor,omitempty"`
		BoardVersion    string `json:"board_version,omitempty"`
		BoardSerial     string `json:"board_serial,omitempty"`
	} `json:"node_dmi_info,omitempty"`
	NodeNetworkInfo map[string]struct {
		Broadcast  string `json:"broadcast,omitempty"`
		Device     string `json:"device,omitempty"`
		Adminstate string `json:"adminstate,omitempty"`
		Operstate  string `json:"operstate,omitempty"`
		Address    string `json:"address,omitempty"`
		Duplex     string `json:"duplex,omitempty"`
	} `json:"node_network_info,omitempty"`
}
