package EDDNClient

// Ring describes planetary rings of a body that may or may not be included
// in a journal message.
type Ring struct {
	InnerRad  float64 `mapstructure:"InnerRad" json:"InnerRad"`
	MassMT    float64 `mapstructure:"MassMT" json:"MassMT"`
	Name      string  `mapstructure:"Name" json:"Name"`
	OuterRad  float64 `mapstructure:"OuterRad" json:"OuterRad"`
	RingClass string  `mapstructure:"RingClass" json:"RingClass"`
}

// Composition describes atmospheric composition that may or may not be
// included in a journal Message.
type Composition struct {
	Name    string  `mapstructure:"Name" json:"Name"`
	Percent float64 `mapstructure:"Percent" json:"Percent"`
}

// Material describes the name, and percentage contained on a planet, or moon.
type Material struct {
	Name    string  `mapstructure:"Name" json:"Name"`
	Percent float64 `mapstructure:"Percent" json:"Percent"`
}

// Faction describes an individual faction that may or may not be included
// in the journal Message.
type Faction struct {
	Allegiance   string  `mapstructure:"Allegiance" json:"Allegiance"`
	FactionState string  `mapstructure:"FactionState" json:"FactionState"`
	Government   string  `mapstructure:"Government" json:"Government"`
	Influence    float32 `mapstructure:"Influence" json:"Influence"`
	Name         string  `mapstructure:"Name" json:"Name"`
}

// JournalDocked contains information pertaining to a 'docked' event.  This
// is missing the 'Security' field, but it seems to mostly go unused with this
// event so it's omitted for now.
type JournalDocked struct {
	StarSystem        string    `mapstructure:"StarSystem" json:"StarSystem"`
	StationFaction    string    `mapstructure:"StationFaction" json:"StationFaction"`
	StationGovernment string    `mapstructure:"StationGovernment" json:"StationGovernment"`
	Timestamp         string    `mapstructure:"timestamp" json:"timestamp"`
	StationAllegiance string    `mapstructure:"StationAllegiance" json:"StationAllegiance"`
	StationEconomy    string    `mapstructure:"StationEconomy" json:"StationEconomy"`
	StarPos           []float64 `mapstructure:"StarPos" json:"StarPos"`
	StationName       string    `mapstructure:"StationName" json:"StationName"`
	StationType       string    `mapstructure:"StationType" json:"StationType"`
	DistFromStarLS    float64   `mapstructure:"DistFromStarLS" json:"DistFromStarLS"`
	FactionState      string    `mapstructure:"FactionState" json:"FactionState"`
	Event             string    `mapstructure:"event" json:"event"`
}

// JournalScanStar contains information about a scanned star.  This is used
// when a journal entry has a StarType field.  Barring that a JournalScanPlanet
// type will be used.
type JournalScanStar struct {
	StellarMass           float64   `mapstructure:"StellarMass" json:"StellarMass"`
	BodyName              string    `mapstructure:"BodyName" json:"BodyName"`
	StarSystem            string    `mapstructure:"StarSystem" json:"StarSystem"`
	Timestamp             string    `mapstructure:"timestamp" json:"timestamp"`
	RotationPeriod        float64   `mapstructure:"RotationPeriod" json:"RotationPeriod"`
	Rings                 []Ring    `mapstructure:"Rings" json:"Rings"`
	StarType              string    `mapstructure:"StarType" json:"StarType"`
	Radius                float64   `mapstructure:"Radius" json:"Radius"`
	AbsoluteMagnitude     float64   `mapstructure:"AbsoluteMagnitude" json:"AbsoluteMagnitude"`
	StarPos               []float64 `mapstructure:"StarPos" json:"StarPos"`
	AgeMy                 int       `mapstructure:"Age_MY" json:"Age_MY"`
	Event                 string    `mapstructure:"event" json:"event"`
	DistanceFromArrivalLS float64   `mapstructure:"DistanceFromArrivalLS" json:"DistanceFromArrivalLS"`
	SurfaceTemperature    float64   `mapstructure:"SurfaceTemperature" json:"SurfaceTemperature"`
	Eccentricity          float64   `mapstructure:"Eccentricity" json:"Eccentricity"`
	OrbitalInclination    float64   `mapstructure:"OrbitalInclination" json:"OrbitalInclination"`
	OrbitalPeriod         float64   `mapstructure:"OrbitalPeriod" json:"OrbitalPeriod"`
	Periapsis             float64   `mapstructure:"Periapsis" json:"Periapsis"`
	SemiMajorAxis         float64   `mapstructure:"SemiMajorAxis" json:"SemiMajorAxis"`
}

// JournalScanPlanet contains information about a scanned moon, or planet.
// This is used when a journal entry does NOT have a StarType field.  If it
// does then a JournalScanStar type will be used.
type JournalScanPlanet struct {
	Eccentricity          float64    `mapstructure:"Eccentricity" json:"Eccentricity"`
	OrbitalInclination    float64    `mapstructure:"OrbitalInclination" json:"OrbitalInclination"`
	OrbitalPeriod         float64    `mapstructure:"OrbitalPeriod" json:"OrbitalPeriod"`
	Periapsis             float64    `mapstructure:"Periapsis" json:"Periapsis"`
	SemiMajorAxis         float64    `mapstructure:"SemiMajorAxis" json:"SemiMajorAxis"`
	BodyName              string     `mapstructure:"BodyName" json:"BodyName"`
	DistanceFromArrivalLS float64    `mapstructure:"DistanceFromArrivalLS" json:"DistanceFromArrivalLS"`
	TidalLock             bool       `mapstructure:"TidalLock" json:"TidalLock"`
	TerraformState        string     `mapstructure:"TerraformState" json:"TerraformState"`
	PlanetClass           string     `mapstructure:"PlanetClass" json:"PlanetClass"`
	SurfacePressure       float64    `mapstructure:"SurfacePressure" json:"SurfacePressure"`
	MassEM                float64    `mapstructure:"MassEM" json:"MassEM"`
	RotationPeriod        float64    `mapstructure:"RotationPeriod" json:"RotationPeriod"`
	Event                 string     `mapstructure:"event" json:"event"`
	StarPos               []float64  `mapstructure:"StarPos" json:"StarPos"`
	AtmosphereType        string     `mapstructure:"AtmosphereType" json:"AtmosphereType"`
	SurfaceTemperature    float64    `mapstructure:"SurfaceTemperature" json:"SurfaceTemperature"`
	Timestamp             string     `mapstructure:"timestamp" json:"timestamp"`
	Materials             []Material `mapstructure:"Materials" json:"Materials"`
	Volcanism             string     `mapstructure:"Volcanism" json:"Volcanism"`
	StarSystem            string     `mapstructure:"StarSystem" json:"StarSystem"`
	Atmosphere            string     `mapstructure:"Atmosphere" json:"Atmosphere"`
	Landable              bool       `mapstructure:"Landable" json:"Landable"`
	Radius                float64    `mapstructure:"Radius" json:"Radius"`
	SurfaceGravity        float64    `mapstructure:"SurfaceGravity" json:"SurfaceGravity"`
}

// JournalFSDJump contains information about a system after a frameshift
// jump is performed.
type JournalFSDJump struct {
	StarSystem       string    `mapstructure:"StarSystem" json:"StarSystem"`
	Timestamp        string    `mapstructure:"timestamp" json:"timestamp"`
	Event            string    `mapstructure:"event" json:"event"`
	SystemSecurity   string    `mapstructure:"SystemSecurity" json:"SystemSecurity"`
	SystemAllegiance string    `mapstructure:"SystemAllegiance" json:"SystemAllegiance"`
	SystemEconomy    string    `mapstructure:"SystemEconomy" json:"SystemEconomy"`
	StarPos          []float64 `mapstructure:"StarPos" json:"StarPos"`
	SystemGovernment string    `mapstructure:"SystemGovernment" json:"SystemGovernment"`
}

// Journal is the high level type that contains the entire JSON message.
type Journal struct {
	SchemaRef string      `json:"$schemaRef"`
	Header    Header      `json:"header"`
	Message   interface{} `json:"message"`
}
