package services

type WinboxEntry struct {
	Crc     int    `json:"crc,omitempty"`
	Version string `json:"version,omitempty"`
	Size    int    `json:"size,omitempty"`
}

type MikrotikWinbox struct {
	Index map[string]WinboxEntry `json:"index"`
	List  map[string]WinboxEntry `json:"list"`
}
