package services

type Domoticz struct {
	DzventsVersion string `json:"dzvents_version,omitempty"`
	Hash           string `json:"hash,omitempty"`
	PythonVersion  string `json:"python_version,omitempty"`
	BuildTime      string `json:"build_time,omitempty"`
}
