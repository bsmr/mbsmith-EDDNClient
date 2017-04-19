package journal

// Header contains basic metadata about the message sent.
type Header struct {
	GatewayTimestamp string `json:"gatewayTimestamp,omitempty"`
	SoftwareName     string `json:"softwareName"`
	SoftwareVersion  string `json:"softwareVersion"`
	UploaderID       string `json:"uploaderID"`
}

// Ring describes planetary rings of a body that may or may not be included
// in a journal message.
type Ring struct {
	InnerRad  float64 `json:"InnerRad"`
	MassMT    float64 `json:"MassMT"`
	Name      string  `json:"Name"`
	OuterRad  float64 `json:"OuterRad"`
	RingClass string  `json:"RingClass"`
}

// Composition describes atmospheric composition that may or may not be
// included in a journal Message.
type Composition struct {
	Name    string  `json:"Name"`
	Percent float64 `json:"Percent"`
}

// Faction describes an individual faction that may or may not be included
// in the journal Message.
type Faction struct {
	Allegiance   string  `json:"Allegiance"`
	FactionState string  `json:"FactionState"`
	Government   string  `json:"Government"`
	Influence    float32 `json:"Influence"`
	Name         string  `json:"Name"`
}

// Message contains the actual data sent to EDDN.
type Message struct {
	AbsoluteMagnitude     float64       `json:"AbsoluteMagnitude"`
	AgeMy                 int           `json:"Age_MY"`
	Atmosphere            string        `json:"Atmosphere"`
	AtmosphereComposition []Composition `json:"AtmosphereComposition"`
	AtmosphereType        string        `json:"AtmosphereType"`
	BodyName              string        `json:"BodyName"`
	DistanceFromArrivalLS float64       `json:"DistanceFromArrivalLS"`
	DistFromStarLS        float64       `json:"DistFromStarLS"`
	Eccentricity          float64       `json:"Eccentricity"`
	Factions              []Faction     `json:"Factions"`
	FactionState          string        `json:"FactionState"`
	Landable              bool          `json:"Landable"`
	MassEM                float64       `json:"MassEM"`
	OrbitalInclination    float64       `json:"OrbitalInclination"`
	OrbitalPeriod         float64       `json:"OrbitalPeriod"`
	Periapsis             float64       `json:"Periapsis"`
	PowerplayState        string        `json:"PowerplayState"`
	Powers                []string      `json:"Powers"`
	Radius                float64       `json:"Radius"`
	Rings                 []Ring        `json:"Rings"`
	RotationPeriod        float64       `json:"RotationPeriod"`
	SemiMajorAxis         float64       `json:"SemiMajorAxis"`
	CockpitBreach         int           `json:"CockpitBreach,omitempty"`
	Event                 string        `json:"event"`
	StarPos               []float64     `json:"StarPos"`
	StarSystem            string        `json:"StarSystem"`
	StarType              string        `json:"StarType"`
	StationEconomy        string        `json:"StationEconomy"`
	StationFaction        string        `json:"StationFaction"`
	StationGovernment     string        `json:"StationGovernment"`
	StationName           string        `json:"StationName"`
	StationType           string        `json:"StationType"`
	SurfaceTemperature    float64       `json:"SurfaceTemperature"`
	SurfaceGravity        float64       `json:"SurfaceGravity"`
	SurfacePressure       float64       `json:"SurfacePressure"`
	SystemAllegiance      string        `json:"SystemAllegiance,omitempty"`
	SystemEconomy         string        `json:"SystemEconomy"`
	SystemGovernment      string        `json:"SystemGovernment"`
	SystemSecurity        string        `json:"SystemSecurity"`
	TerraformState        string        `json:"TerraformState,omitempty"`
	TidalLock             bool          `json:"TidalLock"`
	Timestamp             string        `json:"timestamp"`
	Volcanism             string        `json:"Volcanism"`
}

// Root is the high level type that contains the entire JSON message.
type Root struct {
	SchemaRef string  `json:"$schemaRef"`
	Header    Header  `json:"header"`
	Message   Message `json:"message"`
}
