package EDDNClient

// OutfittingMessage contains the outfitting data sent to EDDN.
type OutfittingMessage struct {
	Modules     []string `json:"modules"`     // Required
	StationName string   `json:"stationName"` // Required
	SystemName  string   `json:"systemName"`  // Required
	Timestamp   string   `json:"timestamp"`   // Required
}

// Outfitting is the high level type that contains the entire JSON message.
type Outfitting struct {
	SchemaRef string            `json:"$schemaRef"`
	Header    Header            `json:"header"`
	Message   OutfittingMessage `json:"message"`
}
