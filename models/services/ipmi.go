package services

type IPMI struct {
	UserAuth     []string `json:"user_auth"`
	Oemid        int      `json:"oemid"`
	Version      string   `json:"version"`
	PasswordAuth []string `json:"password_auth"`
	Level        []string `json:"level"`
}
