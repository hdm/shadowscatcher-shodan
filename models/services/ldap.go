package services

type LDAP struct {
	ResultCode           string   `json:"resultCode,omitempty"`
	ErrorMessage         string   `json:"errorMessage,omitempty"`
	SubschemaSubentry    string   `json:"subschemaSubentry,omitempty"`
	SupportedControl     []string `json:"supportedControl,omitempty"`
	NamingContexts       any      `json:"namingContexts,omitempty"` // TODO: string or []string
	SupportedExtension   []string `json:"supportedExtension,omitempty"`
	SupportedLDAPVersion string   `json:"supportedLDAPVersion,omitempty"`
}
