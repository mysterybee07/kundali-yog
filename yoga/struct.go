package yoga

type ConjunctionType string

type Condition func(chart Chart) bool

const (
	HouseBased    ConjunctionType = "house"
	DegreeBased   ConjunctionType = "degree"
	StrengthBased ConjunctionType = "strength"
)

type PlanetDignity struct {
	OwnSigns       []int
	ExaltationSign int
	Debilitation   int
}

type PlanetData struct {
	PlanetName string  `json:"PlanetEngName"`
	Sign       int     `json:"Sign"`
	House      int     `json:"House"`
	Degree     float64 `json:"Degree"`
	Retro      bool    `json:"Retro"`
	Combust    bool    `json:"Combust"`
	Power      string  `json:"Power,omitempty"`
	Nakshatra  Nakshatra
}

type Nakshatra struct {
	Name string `json:"NakshatraName"`
	Lord string `json:"NakshatraLord"`
	Pada int    `json:"NakshatraPada,string"`
}

type Planet string

const (
	Sun     Planet = "Sun"
	Moon    Planet = "Moon"
	Mars    Planet = "Mars"
	Mercury Planet = "Mercury"
	Jupiter Planet = "Jupiter"
	Venus   Planet = "Venus"
	Saturn  Planet = "Saturn"
	Rahu    Planet = "Rahu"
	Ketu    Planet = "Ketu"
	Asc     Planet = "Asc"
)

type Placement struct {
	Planet  Planet
	House   int
	Sign    int
	Degree  float64
	Retro   bool
	Combust bool
	Power   string
}

type Chart struct {
	Planets map[Planet]Placement
}

type ChartPayload struct {
	Planets []PlanetData `json:"planets"`
}

type YogaStrength string

const (
	YogaStrengthStrong YogaStrength = "strong"
	YogaStrengthWeak   YogaStrength = "weak"
)

type Yoga struct {
	Name       string
	Group      string
	Conditions []Condition
	WeakIf     []Condition
}

type YogaResult struct {
	Name     string       `json:"name"`
	Group    string       `json:"group,omitempty"`
	Strength YogaStrength `json:"strength"`
}

type ConjunctionConfig struct {
	Type ConjunctionType
	Orb  float64
}
