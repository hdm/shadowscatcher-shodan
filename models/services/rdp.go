package services

type RDPEncryption struct {
	Levels    []string `json:"levels"`
	Methods   []string `json:"methods"`
	Protocols []string `json:"protocols"`
}
