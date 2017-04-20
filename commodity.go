package EDDNClient

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

// CommodityMessage contains the commodity data sent to EDDN.
type CommodityMessage struct {
	Commodities []Commodities `json:"commodities"` // Required
	StationName string        `json:"stationName"` // Required
	SystemName  string        `json:"systemName"`  // Required
	Timestamp   string        `json:"timestamp"`   // Required
}

// Commodity is the high level type that contains the entire JSON message.
type Commodity struct {
	SchemaRef string           `json:"$schemaRef"`
	Header    Header           `json:"header"`
	Message   CommodityMessage `json:"message"`
}
