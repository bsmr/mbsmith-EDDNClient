package commodity

// Commodities describes various commodities sent in a Message.
type Commodities struct {
	BuyPrice      int      `json:"buyPrice"`
	Demand        int      `json:"demand"`
	DemandBracket int      `json:"demandBracket"`
	MeanPrice     int      `json:"meanPrice"`
	Name          string   `json:"name"`
	SellPrice     int      `json:"sellPrice"`
	StatusFlags   []string `json:"statusFlags,omitempty"`
	Stock         int      `json:"stock"`
	StockBracket  int      `json:"stockBracket"`
}

// Header contains basic metadata about the message sent.
type Header struct {
	GatewayTimestamp string `json:"gatewayTimestamp,omitempty"`
	SoftwareName     string `json:"softwareName"`
	SoftwareVersion  string `json:"softwareVersion"`
	UploaderID       string `json:"uploaderID"`
}

// Message contains the actual data sent to EDDN.
type Message struct {
	Commodities []Commodities `json:"commodities"` // Required
	StationName string        `json:"stationName"` // Required
	SystemName  string        `json:"systemName"`  // Required
	Timestamp   string        `json:"timestamp"`   // Required
}

// Root is the high level type that contains the entire JSON message.
type Root struct {
	SchemaRef string  `json:"$schemaRef"`
	Header    Header  `json:"header"`
	Message   Message `json:"message"`
}
