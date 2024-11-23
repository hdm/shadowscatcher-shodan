package services

type LDAP struct {
	ResultCode           string   `json:"resultCode,omitempty"`
	ErrorMessage         string   `json:"errorMessage,omitempty"`
	SubschemaSubentry    string   `json:"subschemaSubentry,omitempty"`
	SupportedControl     []string `json:"supportedControl,omitempty"`
	NamingContexts       string   `json:"namingContexts,omitempty"`
	SupportedExtension   []string `json:"supportedExtension,omitempty"`
	SupportedLDAPVersion string   `json:"supportedLDAPVersion,omitempty"`
}
