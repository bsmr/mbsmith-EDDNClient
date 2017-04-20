package EDDNClient

// ShipyardMessage contains the shipyard data sent to EDDN.
type ShipyardMessage struct {
	Ships       []string `json:"ships"`       // Required
	StationName string   `json:"stationName"` // Required
	SystemName  string   `json:"systemName"`  // Required
	Timestamp   string   `json:"timestamp"`   // Required
}

// Shipyard is the high level type that contains the entire JSON message.
type Shipyard struct {
	SchemaRef string          `json:"$schemaRef"`
	Header    Header          `json:"header"`
	Message   ShipyardMessage `json:"message"`
}
