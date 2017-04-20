package blackmarket

// Header contains basic metadata about the message sent.
type Header struct {
	GatewayTimestamp string `json:"gatewayTimestamp,omitempty"`
	SoftwareName     string `json:"softwareName"`
	SoftwareVersion  string `json:"softwareVersion"`
	UploaderID       string `json:"uploaderID"`
}

// Message contains the actual data sent to EDDN.
type Message struct {
	Name        string `json:"name"`        // Required
	Prohibited  bool   `json:"prohibited"`  // Required
	SellPrice   int    `json:"sellPrice"`   // Required
	StationName string `json:"stationName"` // Required
	SystemName  string `json:"systemName"`  // Required
	Timestamp   string `json:"timestamp"`   // Required
}

// Root is the high level type that contains the entire JSON message.
type Root struct {
	SchemaRef string  `json:"$schemaRef"`
	Header    Header  `json:"header"`
	Message   Message `json:"message"`
}
