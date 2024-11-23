package services

type Consul struct {
	Datacenter        string `json:"Datacenter,omitempty"`
	PrimaryDatacenter string `json:"PrimaryDatacenter,omitempty"`
	NodeName          string `json:"NodeName,omitempty"`
	NodeID            string `json:"NodeID,omitempty"`
	Server            bool   `json:"Server,omitempty"`
	Version           string `json:"Version,omitempty"`
	Revision          string `json:"Revision,omitempty"`
}
