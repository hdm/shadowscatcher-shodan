package services

type AFP struct {
	ServerName       string   `json:"server_name,omitempty"`
	NetworkAddresses []string `json:"network_addresses,omitempty"`
	ServerSignature  string   `json:"server_signature,omitempty"`
	Uams             []string `json:"uams,omitempty"`
	Utf8ServerName   string   `json:"utf8_server_name,omitempty"`
	MachineType      string   `json:"machine_type,omitempty"`
	ServerFlags      struct {
		ServerNotifications      bool   `json:"server_notifications,omitempty"`
		Reconnect                bool   `json:"reconnect,omitempty"`
		PasswordSavingProhibited bool   `json:"password_saving_prohibited,omitempty"`
		CopyFile                 bool   `json:"copy_file,omitempty"`
		PasswordChanging         bool   `json:"password_changing,omitempty"`
		ServerSignature          bool   `json:"server_signature,omitempty"`
		Uuids                    bool   `json:"uuids,omitempty"`
		OpenDirectory            bool   `json:"open_directory,omitempty"`
		ServerMessages           bool   `json:"server_messages,omitempty"`
		FlagHex                  string `json:"flag_hex,omitempty"`
		Utf8ServerName           bool   `json:"utf8_server_name,omitempty"`
		SuperClient              bool   `json:"super_client,omitempty"`
		TCPIP                    bool   `json:"tcp_ip,omitempty"`
	} `json:"server_flags,omitempty"`
	AfpVersions    []string `json:"afp_versions,omitempty"`
	DirectoryNames []string `json:"directory_names,omitempty"`
}
