package services

type BGP struct {
	Messages []struct {
		Type          string `json:"type,omitempty"`
		Length        int    `json:"length,omitempty"`
		Version       int    `json:"version,omitempty"`
		BgpIdentifier string `json:"bgp_identifier,omitempty"`
		HoldTime      int    `json:"hold_time,omitempty"`
		Asn           int    `json:"asn,omitempty"`
		ErrorCode     string `json:"error_code,omitempty"`
		ErrorSubcode  string `json:"error_subcode,omitempty"`
	} `json:"messages,omitempty"`
}
