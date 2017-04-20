package EDDNClient

// BlackmarketMessage contains the blackmarket data sent to EDDN.
type BlackmarketMessage struct {
	Name        string `json:"name"`        // Required
	Prohibited  bool   `json:"prohibited"`  // Required
	SellPrice   int    `json:"sellPrice"`   // Required
	StationName string `json:"stationName"` // Required
	SystemName  string `json:"systemName"`  // Required
	Timestamp   string `json:"timestamp"`   // Required
}

// Blackmarket is the high level type that contains the entire JSON message.
type Blackmarket struct {
	SchemaRef string             `json:"$schemaRef"`
	Header    Header             `json:"header"`
	Message   BlackmarketMessage `json:"message"`
}
