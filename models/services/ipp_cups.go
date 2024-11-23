package services

type IPPCups struct {
	StatusMessage string `json:"status_message"`
	Printers      []struct {
		MakeAndModel string `json:"make_and_model,omitempty"`
		URISupported string `json:"uri_supported,omitempty"`
		Name         string `json:"name,omitempty"`
		Location     string `json:"location,omitempty"`
	} `json:"printers,omitempty"`
}
