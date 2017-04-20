package EDDNClient

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

// JournalMessage contains the journal data sent to EDDN.
type JournalMessage struct {
	AbsoluteMagnitude     float64       `json:"AbsoluteMagnitude,omitempty"`
	AgeMy                 int           `json:"Age_MY,omitempty"`
	Atmosphere            string        `json:"Atmosphere,omitempty"`
	AtmosphereComposition []Composition `json:"AtmosphereComposition,omitempty"`
	AtmosphereType        string        `json:"AtmosphereType,omitempty"`
	BodyName              string        `json:"BodyName,omitempty"`
	DistanceFromArrivalLS float64       `json:"DistanceFromArrivalLS,omitempty"`
	DistFromStarLS        float64       `json:"DistFromStarLS,omitempty"`
	Eccentricity          float64       `json:"Eccentricity,omitempty"`
	Factions              []Faction     `json:"Factions,omitempty"`
	FactionState          string        `json:"FactionState,omitempty"`
	Landable              bool          `json:"Landable,omitempty"`
	MassEM                float64       `json:"MassEM,omitempty"`
	OrbitalInclination    float64       `json:"OrbitalInclination,omitempty"`
	OrbitalPeriod         float64       `json:"OrbitalPeriod,omitempty"`
	Periapsis             float64       `json:"Periapsis,omitempty"`
	PowerplayState        string        `json:"PowerplayState,omitempty"`
	Powers                []string      `json:"Powers,omitempty"`
	Radius                float64       `json:"Radius,omitempty"`
	Rings                 []Ring        `json:"Rings,omitempty"`
	RotationPeriod        float64       `json:"RotationPeriod,omitempty"`
	SemiMajorAxis         float64       `json:"SemiMajorAxis,omitempty"`
	Event                 string        `json:"event"`      //Required
	StarPos               []float64     `json:"StarPos"`    //Required
	StarSystem            string        `json:"StarSystem"` // Required
	StarType              string        `json:"StarType,omitempty"`
	StationEconomy        string        `json:"StationEconomy,omitempty"`
	StationFaction        string        `json:"StationFaction,omitempty"`
	StationGovernment     string        `json:"StationGovernment,omitempty"`
	StationName           string        `json:"StationName,omitempty"`
	StationType           string        `json:"StationType,omitempty"`
	SurfaceTemperature    float64       `json:"SurfaceTemperature,omitempty"`
	SurfaceGravity        float64       `json:"SurfaceGravity,omitempty"`
	SurfacePressure       float64       `json:"SurfacePressure,omitempty"`
	SystemAllegiance      string        `json:"SystemAllegiance,omitempty"`
	SystemEconomy         string        `json:"SystemEconomy,omitempty"`
	SystemGovernment      string        `json:"SystemGovernment,omitempty"`
	SystemSecurity        string        `json:"SystemSecurity,omitempty"`
	TerraformState        string        `json:"TerraformState,omitempty"`
	TidalLock             bool          `json:"TidalLock,omitempty"`
	Timestamp             string        `json:"timestamp"` //Required
	Volcanism             string        `json:"Volcanism,omitempty"`
}

// Journal is the high level type that contains the entire JSON message.
type Journal struct {
	SchemaRef string         `json:"$schemaRef"`
	Header    Header         `json:"header"`
	Message   JournalMessage `json:"message"`
}
