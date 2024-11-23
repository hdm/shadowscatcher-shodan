package services

type AMQP struct {
	ServerProperties struct {
		Information  string `json:"information"`
		Product      string `json:"product"`
		Copyright    string `json:"copyright"`
		Capabilities struct {
			ExchangeExchangeBindings   bool `json:"exchange_exchange_bindings"`
			ConnectionBlocked          bool `json:"connection.blocked"`
			AuthenticationFailureClose bool `json:"authentication_failure_close"`
			BasicNack                  bool `json:"basic.nack"`
			PerConsumerQos             bool `json:"per_consumer_qos"`
			ConsumerPriorities         bool `json:"consumer_priorities"`
			ConsumerCancelNotify       bool `json:"consumer_cancel_notify"`
			PublisherConfirms          bool `json:"publisher_confirms"`
			DirectReplyTo              bool `json:"direct_reply_to"`
		} `json:"capabilities"`
		ClusterName string `json:"cluster_name"`
		Platform    string `json:"platform"`
		Version     string `json:"version"`
	} `json:"server_properties"`
	VersionMinor    int      `json:"version_minor"`
	Mechanisms      string   `json:"mechanisms"`
	Locales         string   `json:"locales"`
	VersionMajor    int      `json:"version_major"`
	ProtocolVersion string   `json:"protocol_version"`
	SaslMechanisms  []string `json:"sasl_mechanisms,omitempty"`
}
