package shipyard

// Header contains basic metadata about the message sent.
type Header struct {
	GatewayTimestamp string `json:"gatewayTimestamp,omitempty"`
	SoftwareName     string `json:"softwareName"`
	SoftwareVersion  string `json:"softwareVersion"`
	UploaderID       string `json:"uploaderID"`
}

// Message contains the actual data sent to EDDN.
type Message struct {
	Ships       []string `json:"ships"`       // Required
	StationName string   `json:"stationName"` // Required
	SystemName  string   `json:"systemName"`  // Required
	Timestamp   string   `json:"timestamp"`   // Required
}

// Root is the high level type that contains the entire JSON message.
type Root struct {
	SchemaRef string  `json:"$schemaRef"`
	Header    Header  `json:"header"`
	Message   Message `json:"message"`
}
