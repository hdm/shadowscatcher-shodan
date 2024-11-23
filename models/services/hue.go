package services

type HueBridge struct {
	ModelID          string `json:"model_id,omitempty"`
	Name             string `json:"name,omitempty"`
	SwVersion        string `json:"sw_version,omitempty"`
	BridgeID         string `json:"bridge_id,omitempty"`
	Mac              string `json:"mac,omitempty"`
	FactoryNew       bool   `json:"factory_new,omitempty"`
	DataStoreVersion string `json:"data_store_version,omitempty"`
	APIVersion       string `json:"api_version,omitempty"`
}
