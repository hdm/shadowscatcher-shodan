package services

type CouchDB struct {
	Reason      string `json:"reason,omitempty"`
	HTTPHeaders string `json:"http_headers,omitempty"`
	Error       string `json:"error,omitempty"`
}
