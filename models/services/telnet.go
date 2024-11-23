package services

type Telnet struct {
	Will []string `json:"will,omitempty"`
	Do   []string `json:"do,omitempty"`
	Dont []string `json:"dont,omitempty"`
	Wont []string `json:"wont,omitempty"`
}
